package contract

type CacheInterface interface {
	Get(key string) ConstraintInterface
	Set(key string, value ConstraintInterface)
	Exist(key string) bool
	Delete(key string)
}
