package config

import (
	"encoding/json"
	"errors"
	"github.com/chenxyzl/gsgen/gsmodel"
	"github.com/spf13/viper"
	"os"
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
	err := loadFile()
	if err == nil {
		return nil
	}
	if globalConfig.Load() == nil {
		return errors.Join(err, errors.New("first load config err"))
	} else {
		return errors.Join(err, errors.New("reload config err"))
	}
}

// load config file, will fire OnLoadData event
func loadFile() error {
	//
	viper.Reset()
	//
	viper.SetConfigType("toml") // or viper.SetConfigType("YAML")
	//
	for _, file := range globalConfigFiles {
		fd, err := os.Open(file)
		if err != nil {
			return err
		}
		//noinspection GoUnhandledErrorResult
		defer fd.Close()
		//
		err = viper.MergeConfig(fd)
		if err != nil {
			return err
		}
	}

	s, err := viper.MarshalToString("json")
	if err != nil {
		return err
	}

	conf := &Config{}
	err = json.Unmarshal([]byte(s), conf)
	if err != nil {
		return err
	}
	globalConfig.Store(conf)
	return nil
}
