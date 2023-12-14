package mysql

import (
	"context"
	db "database/sql"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Config struct {
	ReaderSource   string `envconfig:"READER_SOURCE" required:"true"`
	WriterSource   string `envconfig:"WRITER_SOURCE" required:"true"`
	MaxLifeTimeMin int    `envconfig:"MAX_LIFETIME_MIN" default:"3"`
}

type Driver interface {
	dialect.Driver
	Ping(context.Context) error
}

func SetErrLogger(original *zap.Logger) error {
	l := original.With(zap.String("group", "mysql"))
	stdlog, err := zap.NewStdLogAt(l, zap.ErrorLevel)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(mysql.SetLogger(stdlog))
}

func NewDriver(cfg Config) (Driver, error) {
	rd, err := NewSingleDriver(strings.TrimSpace(cfg.ReaderSource), 10, 100, cfg.MaxLifeTimeMin)
	if err != nil {
		return nil, err
	}

	wd, err := NewSingleDriver(strings.TrimSpace(cfg.WriterSource), 10, 100, cfg.MaxLifeTimeMin)
	if err != nil {
		return nil, err
	}

	return &multiDriver{w: wd, r: rd}, nil
}

func NewSingleDriver(source string, maxIdleConns, maxOpenConns, maxLifetimeMin int) (dialect.Driver, error) {
	d, err := sql.Open("mysql", strings.TrimSpace(source))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db := d.DB()
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeMin) * time.Minute)

	return d, errors.WithStack(db.Ping())
}

type multiDriver struct {
	r, w dialect.Driver
}

func (d *multiDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return errors.WithStack(d.r.Query(ctx, query, args, v))
}

func (d *multiDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return errors.WithStack(d.w.Exec(ctx, query, args, v))
}

func (d *multiDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.w.Tx(ctx)
	return tx, errors.WithStack(err)
}

func (d *multiDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	tx, err := d.w.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	return tx, errors.WithStack(err)
}

func (d *multiDriver) Dialect() string {
	return "mysql"
}

func (d *multiDriver) Close() error {
	rerr := d.r.Close()
	werr := d.w.Close()
	if rerr != nil {
		return errors.WithStack(rerr)
	}
	if werr != nil {
		return errors.WithStack(werr)
	}
	return nil
}

func (d *multiDriver) Ping(ctx context.Context) error {
	if err := d.r.(interface{ DB() *db.DB }).DB().PingContext(ctx); err != nil {
		return errors.WithStack(err)
	}
	if err := d.w.(interface{ DB() *db.DB }).DB().PingContext(ctx); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (d *multiDriver) QueryContext(ctx context.Context, query string, args ...any) (*db.Rows, error) {
	return d.r.(interface{ DB() *db.DB }).DB().QueryContext(ctx, query, args...)
}
