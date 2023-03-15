package repository

type Cache interface {
	Set(key string, value string, ttl int) error
	Get(key string) (string, error)
}
