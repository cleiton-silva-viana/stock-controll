package repository

type IRepository[T any] interface {
	Save(entity T) error
	Update(entity T) error
	Delete(UID string) error
	GetByUID(UID string) (T, error)
}
