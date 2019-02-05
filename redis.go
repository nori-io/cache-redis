// Copyright (C) 2018 The Nori Authors info@nori.io
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program; if not, see <http://www.gnu.org/licenses/>.
package main

import (
	"context"
	"time"

	cfg "github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/nori-io/nori-common/meta"
	noriPlugin "github.com/nori-io/nori-common/plugin"

	"github.com/go-redis/redis"
)

type plugin struct {
	instance interfaces.Cache
	address  cfg.String
	password cfg.String
	db       cfg.Int
}

type instance struct {
	client *redis.Client
}

var (
	Plugin plugin
)

func (p *plugin) Init(_ context.Context, configManager cfg.Manager) error {
	m := configManager.Register(p.Meta())
	p.address = m.String("redis.address", "")
	p.password = m.String("redis.password", "")
	p.db = m.Int("redis.database", "")
	return nil
}

func (p *plugin) Instance() interface{} {
	return p.instance
}

func (p plugin) Meta() meta.Meta {
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/cache/redis",
			Version: "1.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0, <2.0",
		},
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name: "Nori: Redis Cache",
		},
		Interface: meta.Cache,
		License: meta.License{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/"},
		Tags: []string{"cache", "redis"},
	}

}

func (p *plugin) Start(_ context.Context, _ noriPlugin.Registry) error {
	if p.instance == nil {

		instance := &instance{
			client: redis.NewClient(&redis.Options{
				Addr:     p.address(),
				Password: p.password(),
				DB:       p.db(),
			}),
		}

		_, err := instance.client.Ping().Result()
		if err != nil {
			instance.client.Close()
			p.instance = nil
			return err
		}

		p.instance = instance
	}
	return nil
}

func (p *plugin) Stop(_ context.Context, _ noriPlugin.Registry) error {
	err := p.instance.(*instance).client.Close()
	p.instance = nil
	p.address = nil
	p.password = nil
	p.db = nil
	return err
}

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
			return []byte{}, interfaces.CacheKeyNotFound
		}
		return []byte{}, err
	}
	return []byte(val), nil
}

func (i instance) Set(key []byte, value []byte, ttl time.Duration) error {
	return i.client.Set(string(key), value, ttl).Err()
}
