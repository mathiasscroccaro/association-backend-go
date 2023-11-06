package orm

import (
	"captcha_example/internal/domain"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

const (
	SQLite      string = "sqlite"
	PostgresSQL string = "postgres"
)

type GORM struct {
	DB *gorm.DB
}

var db GORM

func GetRepository() *GORM {
	return &db
}

func (repository *GORM) Initialize(databaseURL, databaseType string) {
	var err error

	switch databaseType {
	case SQLite:
		repository.DB, err = gorm.Open(sqlite.Open(databaseURL), &gorm.Config{TranslateError: true})
	case PostgresSQL:
		repository.DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{TranslateError: true})
	default:
		log.Fatalf("Unsupported database type %s", databaseType)
	}

	if err != nil {
		log.Fatalf("Cannot connect to %s database: %s", databaseType, err)
	}
	fmt.Printf("We are connected to the %s database\n", databaseURL)
}

func (repository *GORM) Migrate() {
	if repository.DB == nil {
		log.Fatalf("Database must be initialized before migrate")
	}
	repository.DB.AutoMigrate(&domain.PreRegistration{}, &domain.PersonalDocument{}, &domain.MedicalDocument{})
	fmt.Printf("Schemas were migrated to database\n")
}

func (repository *GORM) DeleteAllData() {
	repository.DB.Exec("DELETE FROM pre_registrations")
}
