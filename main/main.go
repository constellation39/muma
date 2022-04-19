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





type SCInstructorCheck struct {
}

// UserInfo 辅导员信息检查
func InstructorCheck(request *muma.Request) bool {
	body, err := request.Get(fmt.Sprintf("web-gateway/instructor/instructor-check?userId=%d"))
	if err != nil {
		logger.Error("UserInfo", zap.String("id", request.ID), zap.Reflect("err", err))
		return false
	}
	scUserInfo := new(SCUserInfo)
	err = json.Unmarshal(body, scUserInfo)
	if err != nil {
		logger.Error("UserInfo", zap.String("id", request.ID), zap.Reflect("err", err))
		return false
	}
	logger.Debug("UserInfo", zap.String("id", request.ID), zap.Reflect("scUserInfo", scUserInfo))
	return true
}

type SCLoginOut struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// LoginOut 登出
func LoginOut(request *muma.Request) bool {
	body, err := request.Get("web-gateway/user/loginOut")
	if err != nil {
		logger.Error("LoginOut", zap.String("id", request.ID), zap.Reflect("err", err))
		return false
	}
	scLoginOut := new(SCLoginOut)
	err = json.Unmarshal(body, scLoginOut)
	if err != nil {
		logger.Error("LoginOut", zap.String("id", request.ID), zap.Reflect("err", err))
		return false
	}
	logger.Debug("LoginOut", zap.String("id", request.ID), zap.Reflect("scLoginOut", scLoginOut))
	return true
}
