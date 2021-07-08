package cache

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/nori-io/interfaces/nori/cache"
)

type Instance struct {
	client redis.UniversalClient
}

type Config struct {
	Address  []string
	Password string
	DB       int
}

func New(conf Config) (*Instance, error) {
	instance := &Instance{
		client: redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    conf.Address,
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
		if err == redis.Nil {
			return []byte{}, cache.CacheKeyNotFound
		}
		return []byte{}, err
	}
	return []byte(val), nil
}

func (i *Instance) Set(key []byte, value []byte, ttl time.Duration) error {
	return i.client.Set(string(key), value, ttl).Err()
}
