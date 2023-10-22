package orm

import "captcha_example/internal/domain"

func (repository *GORM) CreateAndSavePreRegistration(preRegistration domain.PreRegistration) (domain.PreRegistration, error) {
	if result := repository.DB.Create(&preRegistration); result.Error != nil {
		return domain.PreRegistration{}, result.Error
	}
	return preRegistration, nil
}

func (repository *GORM) GetPreRegistrationByDocumentNumber(documentNumber string) (domain.PreRegistration, error) {
	var preRegistration domain.PreRegistration
	if result := repository.DB.Where("cpf = ?", documentNumber).First(&preRegistration); result.Error != nil {
		if result := repository.DB.Where("rg = ?", documentNumber).First(&preRegistration); result.Error != nil {
			return domain.PreRegistration{}, result.Error
		}
	}
	return preRegistration, nil
}
func (repository *GORM) UpdatePreRegistrationById(id uint, preRegistration domain.PreRegistration) (domain.PreRegistration, error) {
	if result := repository.DB.Model(&preRegistration).Where("id = ?", id).Updates(&preRegistration); result.Error != nil {
		return domain.PreRegistration{}, result.Error
	}
	return preRegistration, nil
}

func (repository *GORM) DeletePreRegistrationById(id uint) error {
	return repository.DB.Delete(&domain.PreRegistration{}, id).Error
}
