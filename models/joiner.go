package models

type Joiner struct {
	Id                   int    `json:"id" gorm:"primaryKey"`
	IdentificationNumber int    `json:"identificationNumber"`
	Name                 string `json:"name"`
	LastName             string `json:"lastName"`
	Stack                string `json:"stack"`
	EnglishLevel         string `json:"englishLevel"`
	DomainExperience     string `json:"domainExperience"`
}
