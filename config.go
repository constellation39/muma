package muma

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
)

// Config 只读线程安全
type globalConfig struct {
	Speed      int    `json:"speed"`
	Debug      bool   `json:"debug"`
	Log        int    `json:"log"`
	Host       string `json:"host"`
	UserConfig string `json:"userConfig"`
	TimeOut    int    `json:"timeOut"`
}

type UserConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var GlobalConfig *globalConfig

func init() {
	GlobalConfig = new(globalConfig)
	if err := LoadConfig("config.json", GlobalConfig); err != nil {
		panic(err)
	}
}

func LoadConfig(path string, vTarge interface{}) error {
	return loadConfig(reflect.ValueOf(vTarge), path)
}

func loadConfig(vTarge reflect.Value, path string) error {
	oTarge := vTarge.Type()
	if oTarge.Elem().Kind() != reflect.Struct {
		return errors.New("type of received parameter is not struct")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, vTarge.Interface())
	if err != nil {
		return err
	}
	return nil
}
