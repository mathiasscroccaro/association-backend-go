package orm

import (
	"captcha_example/internal/domain"
	"strings"
)

func (repository *GORM) CreateAndSavePreRegistration(preRegistration domain.PreRegistration) (domain.PreRegistration, domain.IRepositoryError) {
	if result := repository.DB.Create(&preRegistration); result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicated key not allowed") {
			return domain.PreRegistration{}, domain.ErrPreRegisterDocumentDuplicated
		}
		return domain.PreRegistration{}, domain.ErrOtherRepositoryError
	}
	return preRegistration, nil
}

func (repository *GORM) GetPreRegistrationByDocumentNumber(documentNumber string) (domain.PreRegistration, domain.IRepositoryError) {
	var preRegistration domain.PreRegistration
	if result := repository.DB.Where("cpf = ?", documentNumber).First(&preRegistration); result.Error != nil {
		if result := repository.DB.Where("rg = ?", documentNumber).First(&preRegistration); result.Error != nil {
			return domain.PreRegistration{}, domain.ErrPreRegisterNotFound
		}
	}
	return preRegistration, nil
}
func (repository *GORM) UpdatePreRegistrationById(id uint, preRegistration domain.PreRegistration) (domain.PreRegistration, domain.IRepositoryError) {
	if result := repository.DB.First(&domain.PreRegistration{}, id); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return domain.PreRegistration{}, domain.ErrPreRegisterNotFound
		}
		return domain.PreRegistration{}, domain.ErrOtherRepositoryError
	}
	if result := repository.DB.Model(&preRegistration).Where("id = ?", id).Updates(&preRegistration); result.Error != nil {
		if result.Error.Error() == "duplicated key not allowed" {
			return domain.PreRegistration{}, domain.ErrPreRegisterDocumentDuplicated
		}
		return domain.PreRegistration{}, domain.ErrOtherRepositoryError
	}
	return preRegistration, nil
}

func (repository *GORM) DeletePreRegistrationById(id uint) domain.IRepositoryError {
	if result := repository.DB.First(&domain.PreRegistration{}, id); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return domain.ErrPreRegisterNotFound
		}
		return domain.ErrOtherRepositoryError
	}
	return repository.DB.Where("id = ?", id).Delete(&domain.PreRegistration{}).Error
}
