package pre_register

import (
	"captcha_example/internal/domain"
	"captcha_example/internal/repository"
)

func SavePreRegister(repository repository.IRepository, preRegistration domain.PreRegistration) domain.IApiError {
	_, err := repository.CreateAndSavePreRegistration(preRegistration)

	if err.Error() == domain.ErrPreRegisterDocumentDuplicated.Error() {
		return domain.NewApiError("bad request", 400, "document already registered")
	}

	if err.Error() == domain.ErrOtherRepositoryError.Error() {
		return domain.NewApiError("internal server error", 500, "other repository error")
	}

	return nil
}
