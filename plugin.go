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
	"github.com/nori-io/common/v3/config"
	"github.com/nori-io/common/v3/logger"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
	 "github.com/nori-io/interfaces/cache"

	"github.com/go-redis/redis"
)


type service struct {
	instance *instance
	config *pluginConfig
	logger logger.FieldLogger


}

type pluginConfig struct {
	address     string
	password string
	db  int
}

type instance struct {
	client *redis.Client
}

var (
	Plugin plugin.Plugin = &service{}
)


func (p *service) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.logger = log
	p.config.address = config.String("redis.address", "")()
	p.config.password = config.String("redis.password", "")()
	p.config.db = config.Int("redis.database", "")()
	return nil
}

func (p *service) Instance() interface{} {
	return p.instance
}

func (p *service) Meta() meta.Meta {
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/cache/redis",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io",
		},
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name: "Nori: Redis Cache",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Interface: cache.CacheInterface,
		License: []meta.License{
			{
				Title: "GPLv3",
				Type:  "GPLv3",
				URI:   "https://www.gnu.org/licenses/"},
		},
		Links:      nil,
		Repository: meta.Repository{
			Type: "git",
			URI:  "https://github.com/nori-io/cache-redis",
		},
		Tags:       []string{"cache", "redis"},
	}

}

func (p *service) Start(ctx context.Context, registry plugin.Registry) error {
	if p.instance == nil {

		instance := &instance{
			client: redis.NewClient(&redis.Options{
				Addr:     p.config.address,
				Password: p.config.password,
				DB:       p.config.db,
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

func (p *service) Stop(ctx context.Context, registry plugin.Registry) error {
	//err := p.instance.(*instance).client.Close()
	err:=p.instance.client.Close()
	if err!=nil{
		p.logger.Error(err.Error())
	}

	p.instance = nil

	return err
}

