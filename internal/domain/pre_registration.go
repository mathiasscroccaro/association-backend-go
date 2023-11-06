package domain

import (
	"gorm.io/gorm"
)

type PersonalDocument struct {
	gorm.Model
	PreRegistrationId uint
	Base64Image       string `gorm:"unique" json:"base64Image"`
}

type MedicalDocument struct {
	gorm.Model
	PreRegistrationId uint
	Base64Image       string `gorm:"unique" json:"base64Image"`
}

type PreRegistration struct {
	gorm.Model
	FullName       string `json:"fullName"`
	CPF            string `gorm:"unique" json:"cpf"`
	RG             string `gorm:"unique" json:"rg"`
	BirthDate      string `json:"birthDate"`
	GenderIdentity string `json:"genderIdentity"`
	Nationality    string `json:"nationality"`
	MaritalStatus  string `json:"maritalStatus"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	CEP            string `json:"cep"`
	Street         string `json:"street"`
	HouseNumber    string `json:"houseNumber"`
	Complement     string `json:"complement"`
	Neighborhood   string `json:"neighborhood"`
	City           string `json:"city"`
	State          string `json:"state"`

	PersonalDocuments []PersonalDocument `json:"personalDocuments"`
	MedicalDocuments  []MedicalDocument  `json:"medicalDocuments"`
}

type PreRegistrationRequest struct {
	HashedSolution string `json:"hashedSolution"`
	Solution       string `json:"solution"`

	FormData PreRegistration `json:"formData"`
}

type PreRegistrationResponse struct {
	Detail string
}
