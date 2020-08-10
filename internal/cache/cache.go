package cache

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type Instance struct {
	client *redis.Client
}

type Config struct {
	Address  string
	Password string
	DB       int
}

func New(conf *Config) (*Instance, error) {
	instance := &Instance{
		client: redis.NewClient(&redis.Options{
			Addr:     conf.Address,
			Password: conf.Password,
			DB:       conf.DB,
		}),
	}
	_, err := instance.client.Ping().Result()
	if err != nil {
		instance.client.Close()

	}
	return instance, err
}

func (i *Instance) Close() error {
	return i.client.Close()
}

func (i *Instance) Clear() error {
	return i.client.FlushAll().Err()
}

func (i *Instance) Delete(key []byte) error {
	return i.client.Del(string(key)).Err()
}

func (i *Instance) Get(key []byte) ([]byte, error) {
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

func (i *Instance) Set(key []byte, value []byte, ttl time.Duration) error {
	return i.client.Set(string(key), value, ttl).Err()
}
