package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
	"github.com/kyfk/golang-todo-list-sample/pkg/ent/migrate"
)

var url = flag.String("url", "root:@tcp(127.0.0.1:3306)/todo_dev", "connecting database url")
var service = flag.String("service", "todo", "the service to migrate")

func main() {
	flag.Parse()

	ctx := context.Background()

	client, err := ent.Open("mysql", *url)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	dir, err := sqltool.NewGolangMigrateDir(path.Join("migrations", *service))
	if err != nil {
		log.Fatalf("failed creating golang-migration directory: %v", err)
	}
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(false),
	}

	err = client.Schema.NamedDiff(ctx, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
