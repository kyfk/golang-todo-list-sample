package env

type Env string

const (
	PROD  Env = "prod"
	STG   Env = "stg"
	QA    Env = "qa"
	DEV   Env = "dev"
	LOCAL Env = "local"
)

func (e Env) Is(ee Env) bool {
	return e == ee
}

func (e Env) IsProd() bool {
	return e == PROD
}

func (e Env) IsStg() bool {
	return e == STG
}

func (e Env) IsQA() bool {
	return e == QA
}

func (e Env) IsDev() bool {
	return e == DEV
}

func (e Env) IsLocal() bool {
	return e == LOCAL
}
