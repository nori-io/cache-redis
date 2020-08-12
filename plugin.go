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

	rediscache "github.com/nori-io/cache-redis/internal/cache"
	"github.com/nori-io/common/v3/config"
	"github.com/nori-io/common/v3/logger"
	"github.com/nori-io/common/v3/meta"
	"github.com/nori-io/common/v3/plugin"
	"github.com/nori-io/interfaces/nori/cache"
)

type service struct {
	instance *rediscache.Instance
	config   *pluginConfig
	logger   logger.FieldLogger
}

type pluginConfig struct {
	addresses []string
	password  string
	db        int
}

var (
	Plugin plugin.Plugin = &service{}
)

func (p *service) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.logger = log
	p.config.addresses = config.SliceString("cache.redis.address", "")()
	p.config.password = config.String("cache.redis.password", "")()
	p.config.db = config.Int("cache.redis.database", "")()
	return nil
}

func (p *service) Instance() interface{} {
	return p.instance
}

func (p *service) Meta() meta.Meta {
	return meta.Data{
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
			Name:        "Nori: Redis Cache",
			Description: "",
		},
		Core: meta.Core{
			VersionConstraint: "^0.2.0",
		},
		Interface: cache.CacheInterface,
		License: []meta.License{
			{
				Title: "GPLv3",
				Type:  "GPLv3",
				URI:   "https://www.gnu.org/licenses/"},
		},
		Links: []meta.Link{},
		Repository: meta.Repository{
			Type: "git",
			URI:  "https://github.com/nori-io/cache-redis",
		},
		Tags: []string{"cache", "redis", "cache-redis"},
	}
}

func (p *service) Start(ctx context.Context, registry plugin.Registry) error {

	if p.instance == nil {

		instance, err := rediscache.New(rediscache.Config{
			Address:  p.config.addresses,
			Password: p.config.password,
			DB:       p.config.db,
		})

		if err == nil {
			p.instance = instance
		} else {
			return err
		}
	}
	return nil
}

func (p *service) Stop(ctx context.Context, registry plugin.Registry) error {

	err := p.instance.Close()
	if err != nil {
		p.logger.Error(err.Error())
	}

	p.instance = nil

	return err
}
