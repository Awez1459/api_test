package kv

import "api_test/storage"

type KV interface {
	Set(key string, value *storage.User) error
	Get(key string) (*storage.User, error)
	Update(key string, value *storage.User) error
	GetAll() ([]*storage.User, error)
	Delete(key string) error
}

var inst KV

func Init(store KV) {
	inst = store
}

func Set(key string, value *storage.User) error {
	return inst.Set(key, value)
}

func Get(key string) (*storage.User, error) {
	return inst.Get(key)
}

func Update(key string, value *storage.User) error {
    return inst.Update(key, value)
}

func GetAll() ([]*storage.User, error) {
    return inst.GetAll()
}

func Delete(key string) error {
	return inst.Delete(key)
}
