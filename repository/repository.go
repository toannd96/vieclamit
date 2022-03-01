package repository

import "vieclamit/models"

// Repository interface
type Repository interface {
	Insert(recruitment models.Recruitment) error
	Delete() (int64, error)
	FindByUrl(urlJob string) (int64, error)
	FindByLocation(location string) (*models.Recruitments, error)
	FindBySkill(skill string) (*models.Recruitments, error)
	FindByCompany(company string) (*models.Recruitments, error)
	FindBySkillAndLocation(skill, location string) (*models.Recruitments, error)
	FindByCompanyAndLocation(company, location string) (*models.Recruitments, error)
}
