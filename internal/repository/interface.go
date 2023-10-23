package repository

import "captcha_example/internal/domain"

type IRepository interface {
	Initialize(databaseURL, databaseType string)
	Migrate()
	DeleteAllData()

	CreateAndSavePreRegistration(preRegistration domain.PreRegistration) (domain.PreRegistration, domain.IRepositoryError)
	GetPreRegistrationByDocumentNumber(documentNumber string) (domain.PreRegistration, domain.IRepositoryError)
	UpdatePreRegistrationById(id uint, preRegistration domain.PreRegistration) (domain.PreRegistration, domain.IRepositoryError)
	DeletePreRegistrationById(id uint) domain.IRepositoryError
}
