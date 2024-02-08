package cache

import (
	"github.com/lazygophers/log"
	"time"
)

type fiberStorage interface {
	Reset() error
}

type FiberStorage struct {
	storage Cache
}

func (p *FiberStorage) Get(key string) ([]byte, error) {
	value, err := p.storage.Get(key)
	if err != nil {
		if err == NotFound {
			return nil, nil
		}
		return nil, err
	}
	return []byte(value), nil
}

func (p *FiberStorage) Set(key string, val []byte, exp time.Duration) error {
	return p.storage.SetEx(key, val, exp)
}

func (p *FiberStorage) Delete(key string) error {
	return p.storage.Del(key)
}

func (p *FiberStorage) Reset() error {
	log.Warnf("unsupported operation")
	return nil
}

func (p *FiberStorage) Close() error {
	return p.storage.Close()
}

func GetFiberStorage(storage Cache) fiber.Storage {
	return &FiberStorage{
		storage: storage,
	}
}
