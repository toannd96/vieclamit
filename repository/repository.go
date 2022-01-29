package repository

import "vieclamit/models"

type Repository interface {
	Insert(recruitment models.Recruitment, collection string) error
	Delete(collection string) (int, error)
	FindByUrl(urlJob string, collection string) (int, error)
}
