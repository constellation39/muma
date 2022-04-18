package muma

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

func Load(vTarge reflect.Value, path, name string) error {
	oTarge := vTarge.Type()
	if oTarge.Kind() != reflect.Struct {
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
