package orm

import (
	"captcha_example/internal/domain"
	"testing"
)

func CreateAndSavePreRegistrationTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	preRegistrationData := domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678900",
		RG:             "12345678900",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	preRegistration, err := repository.CreateAndSavePreRegistration(preRegistrationData)

	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	if preRegistration.ID == 0 {
		t.Errorf("Failed to create pre registration: %v", preRegistration)
	}
	if preRegistration.FullName != "John Doe" {
		t.Errorf("Failed to create pre registration. FullName is not expected: %v", preRegistration)
	}
}

func CreateAndSavePreRegistrationWithSameCPFTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	preRegistrationData := domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678900",
		RG:             "12345678900",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	_, err := repository.CreateAndSavePreRegistration(preRegistrationData)
	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	_, err = repository.CreateAndSavePreRegistration(preRegistrationData)
	if err.Error() != "duplicated key not allowed" {
		t.Errorf("Pre Registration with same CPF already returned not expected error: %v", err)
	}
}

func GetPreRegistrationByDocumentNumberTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	preRegistrationData := domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678900",
		RG:             "12345678901",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	_, err := repository.CreateAndSavePreRegistration(preRegistrationData)

	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	preRegistrationByCPF, err := repository.GetPreRegistrationByDocumentNumber("12345678900")

	if err != nil {
		t.Errorf("Failed to get pre registration: %v", err)
	}
	if preRegistrationByCPF.FullName != "John Doe" {
		t.Errorf("Failed to get pre registration: %v", preRegistrationByCPF)
	}
	if preRegistrationByCPF.CPF != "12345678900" {
		t.Errorf("Failed to get pre registration: %v", preRegistrationByCPF)
	}

	preRegistrationByRG, err := repository.GetPreRegistrationByDocumentNumber("12345678901")

	if err != nil {
		t.Errorf("Failed to get pre registration: %v", err)
	}

	if preRegistrationByRG.FullName != preRegistrationByCPF.FullName {
		t.Errorf("PreRegistrationByRG and PreRegistrationByCPF are not equal. By CPF: %v, By RG: %v", preRegistrationByCPF, preRegistrationByRG)
	}
}

func GetPreRegistrationByDocumentNumberNonExistentTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	_, err := repository.GetPreRegistrationByDocumentNumber("12345678900")

	if err.Error() != "record not found" {
		t.Errorf("Error message is not expected: %v", err)
	}
}

func UpdatePreRegistrationByIdTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	preRegistrationData := domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678900",
		RG:             "12345678901",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	preRegistration, err := repository.CreateAndSavePreRegistration(preRegistrationData)

	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	preRegistrationData.FullName = "John Doe 2"
	preRegistrationData.CPF = "12345678901"

	modifiedPreRegistration, err := repository.UpdatePreRegistrationById(preRegistration.ID, preRegistrationData)

	if err != nil {
		t.Errorf("Failed to update pre registration: %v", err)
	}

	if modifiedPreRegistration.FullName != "John Doe 2" {
		t.Errorf("Failed to update pre registration: %v", modifiedPreRegistration)
	}
	if modifiedPreRegistration.CPF != "12345678901" {
		t.Errorf("Failed to update pre registration: %v", modifiedPreRegistration)
	}

}

func UpdatePreRegistrationByIdWithNotAllowedCPFTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	preRegistrationData := domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678900",
		RG:             "12345678901",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	_, err := repository.CreateAndSavePreRegistration(preRegistrationData)

	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	preRegistrationData = domain.PreRegistration{
		FullName:       "John Doe 2",
		CPF:            "22345678900",
		RG:             "22345678901",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	preRegistration, err := repository.CreateAndSavePreRegistration(preRegistrationData)

	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	preRegistrationData.CPF = "12345678900"

	_, err = repository.UpdatePreRegistrationById(preRegistration.ID, preRegistrationData)

	if err.Error() != "duplicated key not allowed" {
		t.Errorf("Failed to update pre registration: %v", err)
	}
}

func UpdatePreRegistrationByIdWithNonExistentTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	_, err := repository.UpdatePreRegistrationById(1, domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678900",
		RG:             "12345678901",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	})

	if err.Error() != "record not found" {
		t.Errorf("Error message is not expected: %v", err)
	}
}

func DeletePreRegistrationByIdTestCase(t *testing.T) {
	tearDownTestCase := setupTestCase(t)
	defer tearDownTestCase(t)

	repository := GetRepository()

	preRegistrationData := domain.PreRegistration{
		FullName:       "John Doe",
		CPF:            "12345678907",
		RG:             "12345678901",
		BirthDate:      "2000-01-01",
		GenderIdentity: "Male",
		Nationality:    "Brazilian",
		MaritalStatus:  "Single",
		Email:          "johndoe@me.com",
		PhoneNumber:    "12345678900",
		CEP:            "12345678",
		Street:         "Street",
		HouseNumber:    "123",
		Complement:     "Complement",
		Neighborhood:   "Neighborhood",
		City:           "City",
		State:          "State",
	}

	preRegistration, err := repository.CreateAndSavePreRegistration(preRegistrationData)

	if err != nil {
		t.Errorf("Failed to create pre registration: %v", err)
	}

	err = repository.DeletePreRegistrationById(preRegistration.ID)

	if err != nil {
		t.Errorf("Failed to delete pre registration: %v", err)
	}
}

func TestPropertyOperations(t *testing.T) {
	setupBeforeAllTestCases(t)

	t.Run("Create a PreRegistration", CreateAndSavePreRegistrationTestCase)
	t.Run("Create a duplicated PreRegistration", CreateAndSavePreRegistrationWithSameCPFTestCase)

	t.Run("Get PreRegistration by Document Number", GetPreRegistrationByDocumentNumberTestCase)
	t.Run("Get PreRegistration by Document Number Non Existent", GetPreRegistrationByDocumentNumberNonExistentTestCase)

	t.Run("Update PreRegistration by ID", UpdatePreRegistrationByIdTestCase)
	t.Run("Update PreRegistration by ID with not allowed CPF", UpdatePreRegistrationByIdWithNotAllowedCPFTestCase)
	t.Run("Update PreRegistration by ID with non existant register", UpdatePreRegistrationByIdWithNonExistentTestCase)

	t.Run("Delete PreRegistration by ID", DeletePreRegistrationByIdTestCase)
}
