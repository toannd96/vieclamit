package repository

type Repository interface {
	Insert(data interface{}, collection string) error
	FindByUrl(urlJob string, collection string) (int, error)
}
