package muma

import (
	"log"
	"testing"
)

func TestLoad(t *testing.T) {
	config := new(Config)
	LoadConfig(config)
	log.Printf("%+v", config)
}
