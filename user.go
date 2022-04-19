package muma

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type User struct {
	*Request
	*Config

	ID       string
	Token    string
	UserId   int
	Phone    int
	CourseId int
}

func NewUser(config *Config) *User {
	return &User{
		Request: NewRequest(config.Host),
		Config:  config,
		ID:      uuid.NewString(),
	}
}

type CSLogin struct {
	Equipment LoginEquipment `json:"equipment"`
	UserName  string         `json:"userName"`
	Password  string         `json:"password"`
}

type LoginEquipment struct {
	DeviceBrand            string `json:"deviceBrand"`
	DeviceModel            string `json:"deviceModel"`
	OperationSystem        string `json:"operationSystem"`
	OperationSystemVersion string `json:"operationSystemVersion"`
	ResolutionHeight       int    `json:"resolutionHeight"`
	ResolutionWidth        int    `json:"resolutionWidth"`
	ServiceProvider        string `json:"serviceProvider"`
}

// Login 登录
func Login(user *User, username, passwrod string) bool {
	Logger.Debug("Login", zap.String("id", user.ID), zap.String("username", username), zap.String("password", passwrod))
	body, err := user.Post("web-gateway/token", CSLogin{
		Equipment: LoginEquipment{
			DeviceBrand:            "Chrome",
			DeviceModel:            fmt.Sprintf("100.0.%d.127", 4000+rand.Intn(2000)),
			OperationSystem:        "Windows 10",
			OperationSystemVersion: "10.0",
			ResolutionHeight:       1920,
			ResolutionWidth:        1080,
			ServiceProvider:        "",
		},
		UserName: username,
		Password: passwrod,
	})
	if err != nil {
		Logger.Error("Login", zap.String("id", user.ID), zap.Reflect("err", err))
		return false
	}
	token := fmt.Sprintf("Bearer %s", body)
	user.Token = token
	user.SetHeader("authorization", token)
	return true
}

type SCRecentlyWork struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data SCRecentlyWorkData `json:"data"`
}
type SCRecentlyWorkData struct {
	Time       int64       `json:"time"`
	ClassName  string      `json:"className"`
	CourseName interface{} `json:"courseName"`
}

// RecentlyWork 最近工作
func RecentlyWork(user *User) bool {
	body, err := user.Get("web-gateway/t/student/recently-work")
	if err != nil {
		Logger.Error("RecentlyWork", zap.String("id", user.ID), zap.Reflect("err", err))
		return false
	}
	scRecentlyWork := new(SCRecentlyWork)
	err = json.Unmarshal(body, scRecentlyWork)
	if err != nil {
		Logger.Error("RecentlyWork", zap.String("id", user.ID), zap.Reflect("err", err))
		return false
	}
	Logger.Debug("RecentlyWork", zap.String("id", user.ID), zap.Reflect("scRecentlyWork", scRecentlyWork))
	return true
}

type SCUserInfo struct {
}

// UserInfo 用户信息
func UserInfo(user *User) bool {
	body, err := user.Get("web-gateway/user/userInfo")
	if err != nil {
		Logger.Error("UserInfo", zap.String("id", user.ID), zap.Reflect("err", err))
		return false
	}
	scUserInfo := new(SCUserInfo)
	err = json.Unmarshal(body, scUserInfo)
	if err != nil {
		Logger.Error("UserInfo", zap.String("id", user.ID), zap.Reflect("err", err))
		return false
	}
	Logger.Debug("UserInfo", zap.String("id", user.ID), zap.Reflect("scUserInfo", scUserInfo))
	return true
}
