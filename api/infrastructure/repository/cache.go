package repository

import (
	"api/domain/repository"
)

type cache struct {
}

func (c cache) Set(key string, value string, ttl int) error {
	//TODO implement me
	return nil
}

func (c cache) Get(key string) (string, error) {
	//TODO implement me
	return "", nil
}

func NewCache() repository.Cache {
	return &cache{}
}
