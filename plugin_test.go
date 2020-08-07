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
/*
import (
	"testing"
	"time"

	"github.com/cheebo/go-config"
	"github.com/nori-io/nori-common/interfaces"
	cfg "github.com/nori-io/nori/core/config"
	"github.com/stretchr/testify/assert"
)

const (
	testRedisAddr = "localhost:6379"
)

var (
	testKey   = []byte("testkey")
	testValue = []byte("testvalue")
)

func TestPackage(t *testing.T) {
	assert := assert.New(t)

	cfgTest := go_config.New()
	cfgTest.SetDefault("redis.address", testRedisAddr)
	cfgTest.SetDefault("redis.password", "")
	cfgTest.SetDefault("redis.database", 0)

	m := cfg.NewManager(cfgTest)

	p := new(plugin)

	assert.NotNil(p.Meta())
	assert.NotEmpty(p.Meta().Id())

	err := p.Init(nil, m)
	assert.Nil(err)

	err = p.Start(nil, nil)
	assert.Nil(err)

	cache, ok := p.Instance().(interfaces.Cache)
	assert.True(ok)
	assert.NotNil(cache)

	err = cache.Set(testKey, testValue, time.Duration(0))
	assert.Nil(err)

	bs, err := cache.Get(testKey)
	assert.Nil(err)
	assert.Equal(testValue, bs)

	err = cache.Delete(testKey)
	assert.Nil(err)

	bs, err = cache.Get(testKey)
	assert.Empty(bs)
	assert.Equal(err, interfaces.CacheKeyNotFound)

	err = cache.Set(testKey, testValue, time.Duration(0))
	assert.Nil(err)

	err = cache.Clear()
	assert.Nil(err)

	bs, err = cache.Get(testKey)
	assert.Empty(bs)
	assert.Equal(err, interfaces.CacheKeyNotFound)

	err = p.Stop(nil, nil)
	assert.Nil(err)
	assert.Nil(p.Instance())
}*/