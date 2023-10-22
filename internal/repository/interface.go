package repository

import "captcha_example/internal/domain"

type IRepository interface {
	Initialize(databaseURL, databaseType string)
	Migrate()
	DeleteAllData()

	CreateAndSavePreRegistration(preRegistration domain.PreRegistration) (domain.PreRegistration, error)
	GetPreRegistrationByDocumentNumber(documentNumber string) (domain.PreRegistration, error)
	UpdatePreRegistrationById(id uint, preRegistration domain.PreRegistration) error
	DeletePreRegistrationById(id uint) error
}
