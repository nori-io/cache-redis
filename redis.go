package main

import (
	"errors"
	"github.com/nori-io/interfaces/cache"
	"time"
)

func (i instance) Clear() error {
	return i.client.FlushAll().Err()
}

func (i instance) Delete(key []byte) error {
	return i.client.Del(string(key)).Err()
}

func (i instance) Get(key []byte) ([]byte, error) {
	val, err := i.client.Get(string(key)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			//return []byte{}, cache.CacheKeyNotFound
			return []byte{}, errors.New("CacheKeyNotFound")

		}
		return []byte{}, err
	}
	return []byte(val), nil
}

func (i instance) Set(key []byte, value []byte, ttl time.Duration) error {
	return i.client.Set(string(key), value, ttl).Err()
}

