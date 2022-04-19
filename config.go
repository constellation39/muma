package muma

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

// Config 只读线程安全
type Config struct {
	Debug    bool   `json:"debug"`
	Log      int    `json:"log"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	TimeOut  int    `json:"timeOut"`
}

var config *Config

func init() {
	config = new(Config)
	LoadConfig(config)
}

func LoadConfig(vTarge interface{}) {
	loadConfig(reflect.ValueOf(vTarge), "", "config.json")
}

func loadConfig(vTarge reflect.Value, path, name string) error {
	oTarge := vTarge.Type()
	if oTarge.Elem().Kind() != reflect.Struct {
		return errors.New("type of received parameter is not struct")
	}
	data, err := ioutil.ReadFile(filepath.Join(path, name))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, vTarge.Interface())
	if err != nil {
		return err
	}
	return nil
}
