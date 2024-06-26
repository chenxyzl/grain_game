package config

import (
	"github.com/chenxyzl/gsgen/gsmodel"
	config_base "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"sync/atomic"
)

type Gate struct {
	wsPath   string `json:"wsPath"`
	wsSubKey string `json:"wsSubKey"`
}

type Config struct {
	app     string                 `toml:"app"`
	version string                 `json:"version"`
	etcd    *gsmodel.AList[string] `json:"etcd"`
	gate    *Gate                  `json:"gate"`
}

var (
	globalConfig      atomic.Value //配置数据
	globalConfigFiles []string
)

func Get() *Config {
	return globalConfig.Load().(*Config)
}

func Init(configFiles ...string) {
	globalConfigFiles = configFiles
}

func Reload() error {
	config_base.ClearAll()
	config_base.WithOptions(config_base.ParseEnv)
	config_base.AddDriver(toml.Driver)
	err := config_base.LoadFiles(globalConfigFiles...)
	if err != nil {
		return err
	}
	var conf = &Config{}
	err = config_base.Decode(conf)
	if err != nil {
		return err
	}
	globalConfig.Store(conf)
	return nil
}
