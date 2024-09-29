package repository

type IRepository[T any] interface {
	Save(entity T) error
	Update(entity T) error
	Delete(ID int) error
	GetByID(ID int) (T, error)
}
