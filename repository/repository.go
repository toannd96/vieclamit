package repository

import "vieclamit/models"

// Repository interface
type Repository interface {
	Insert(recruitment models.Recruitment, collection string) error
	Delete(collection string) (int, error)
	FindByUrl(urlJob, collection string) (int, error)
	FindByLocation(location, collection string) (*models.Recruitments, error)
	FindBySkill(skill, collection string) (*models.Recruitments, error)
	FindByCompany(company, collection string) (*models.Recruitments, error)
}
