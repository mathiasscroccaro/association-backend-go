package repository

import (
	"captcha_example/internal/repository"
	"captcha_example/internal/repository/orm"
	"fmt"
	"os"
)

func GetDB() repository.IRepository {
	return orm.GetRepository()
}

func isThereEnvConfigForPostgres() bool {
	requiredVariables := []string{
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB_NAME",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
	}
	for _, envVar := range requiredVariables {
		if _, isValid := os.LookupEnv(envVar); isValid == false {
			return false
		}
	}
	return true
}

func isThereEnvConfigForSQLite() bool {
	_, isValid := os.LookupEnv("SQLITE_DATABASE_PATH")
	return isValid
}

func getDbURLAndType() (string, string) {
	if isThereEnvConfigForPostgres() {
		postgresUser := os.Getenv("POSTGRES_USER")
		postgresPassword := os.Getenv("POSTGRES_PASSWORD")
		postgresDbName := os.Getenv("POSTGRES_DB_NAME")
		postgresHost := os.Getenv("POSTGRES_HOST")
		postgresPort := os.Getenv("POSTGRES_PORT")
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			postgresUser,
			postgresPassword,
			postgresHost,
			postgresPort,
			postgresDbName,
		), "postgres"
	} else if isThereEnvConfigForSQLite() {
		return os.Getenv("SQLITE_DATABASE_PATH"), "sqlite"
	} else {
		return "./db.sqlite", "sqlite"
	}
}

func InitDBConnectionByEnvConfig() {
	databaseURL, databaseType := getDbURLAndType()

	repository := GetDB()

	repository.Initialize(databaseURL, databaseType)
	repository.Migrate()
	repository.DeleteAllData()
}
