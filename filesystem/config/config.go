package config

import (
	"github.com/deatil/go-filesystem/filesystem/interfaces"
)

/**
 * Config
 *
 * @create 2021-8-1
 * @author deatil
 */
type Config struct {
	data map[string]any
}

func New(data map[string]any) Config {
	return Config{
		data: data,
	}
}

func (this Config) With(data map[string]any) interfaces.Config {
	this.data = data

	return this
}

func (this Config) Set(key string, value any) interfaces.Config {
	this.data[key] = value

	return this
}

func (this Config) Has(key string) bool {
	if _, ok := this.data[key]; ok {
		return true
	}

	return false
}

func (this Config) Get(key string, defaults ...any) any {
	if data, ok := this.data[key]; ok {
		return data
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return nil
}
