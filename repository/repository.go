package repository

type Repository interface {
	Insert(data interface{}, collection string) error
	FindByUrl(url string, collection string) (int, error)
}
