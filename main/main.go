package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"muma"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = muma.Logger
}

func main() {
	request := muma.NewRequest("https://muma.com")

	Login(request, "17780508348", "miku39")
	RecentlyWork(request)
	LoginOut(request)
}







