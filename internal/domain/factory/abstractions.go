package factory

type IFactory[T any, Q any] interface {
	Create(DTO T) (*Q, error)
}
