package interfaces

type GenericRepository[T any] interface {
	Create(data *T) (*T, error)
	Delete(model T, id uint64) error
}
