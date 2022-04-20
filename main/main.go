package main

import (
	"muma"

	"go.uber.org/zap"
)

var users map[string]*muma.User
var names map[string]string
var logger *zap.Logger

func init() {
	logger = muma.Logger
	users = make(map[string]*muma.User)
	names = make(map[string]string)
}

func main() {
	// ctx := context.Background()
	// ReadConfig(ctx)
	// go WatchConfig(ctx)
	user, err := muma.NewUser("user/tanghangran.json")
	if err != nil {
		panic(err)
	}
	logger.Debug("user", zap.Reflect("user", user))

	user.Login()
	scCourseRecordInfo, _ := user.CourseRecordInfo()

	for _, course := range scCourseRecordInfo.List {

	}

	logger.Debug("scCourseRecordInfo", zap.Reflect("scCourseRecordInfo", scCourseRecordInfo))

	defer user.LoginOut()
}

// func Loop(ctx context.Context, user *muma.User) {
// 	timer := time.NewTicker(time.Second * 5)
// 	for {
// 		select {
// 		case <-timer.C:
// 			Run(user)
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// 	user.LoginOut()
// }

// func Run(user *muma.User) {
// 	switch user.State {
// 	case muma.Ready:
// 		user.Login()
// 	case muma.Idle:
// 		user.CourseRecordInfo()
// 	case muma.Look:
// 		// user.LearnedVideo()
// 	}
// }

// func NewUser(path string) error {
// 	user, err := muma.NewUser(path)
// 	if err != nil {
// 		logger.Error("NewUser Failed", zap.String("path", path), zap.Reflect("Error", err))
// 		return err
// 	}
// 	if sou, ok := names[user.Config.Username]; ok {
// 		rUser, ok := users[sou]
// 		if !ok {
// 			delete(names, user.Config.Username)
// 			logger.Error("Replace User Config Failed", zap.String("path", path), zap.String("username", user.Config.Username), zap.String("exits", sou))
// 			return fmt.Errorf("Replace User Config Failed")
// 		}
// 		delete(users, sou)
// 		users[path] = rUser
// 		names[rUser.Config.Username] = path
// 		logger.Warn("User Replace Success", zap.String("replace", path), zap.String("username", user.Config.Username), zap.String("exits", sou))
// 		return fmt.Errorf("User Replace Success")
// 	}
// 	users[path] = user
// 	names[user.Config.Username] = path
// 	logger.Info("NewUser", user.Fields()...)
// 	go Loop(context.Background(), user)
// 	return nil
// }

// func DelUser(path string) error {
// 	user, ok := users[path]
// 	if !ok {
// 		logger.Error("DelUser Failed", zap.String("path", path))
// 		return nil
// 	}
// 	delete(names, user.Config.Username)
// 	delete(users, path)
// 	return nil
// }

// func ReadConfig(ctx context.Context) {
// 	muma.ExitsFile(muma.GlobalConfig.UserConfig)
// 	filepath.Walk(muma.GlobalConfig.UserConfig, func(path string, info fs.FileInfo, err error) error {
// 		if info.IsDir() {
// 			return nil
// 		}
// 		if err != nil {
// 			logger.Error("ReadConfig", zap.String("path", path), zap.Reflect("Error", err))
// 			return err
// 		}
// 		NewUser(path)
// 		return nil
// 	})
// }

// func WatchConfig(ctx context.Context) {
// 	watcher, err := fsnotify.NewWatcher()

// 	if err != nil {
// 		panic(err)
// 	}
// 	muma.ExitsFile(muma.GlobalConfig.UserConfig)
// 	err = watcher.Add(muma.GlobalConfig.UserConfig)

// 	if err != nil {
// 		panic(err)
// 	}

// 	for {
// 		select {
// 		case event, ok := <-watcher.Events:
// 			if !ok {
// 				return
// 			}
// 			if event.Op&fsnotify.Create == fsnotify.Create {
// 				logger.Debug("Create UserConfig File Success", zap.Reflect("event", event))
// 				NewUser(event.Name)
// 			}
// 			if event.Op&fsnotify.Write == fsnotify.Write {
// 				logger.Debug("Change UserConfig File Success", zap.Reflect("event", event))
// 				user, ok := users[event.Name]
// 				if ok {
// 					muma.LoadConfig(event.Name, user.Config)
// 				}
// 			}
// 			if event.Op&fsnotify.Remove == fsnotify.Remove {
// 				logger.Debug("Delete UserConfig File Success", zap.Reflect("event", event))
// 				DelUser(event.Name)
// 			}
// 			if event.Op&fsnotify.Rename == fsnotify.Rename {
// 				logger.Debug("Remove UserConfig File Success", zap.Reflect("event", event))
// 				NewUser(event.Name)
// 			}
// 		case err, ok := <-watcher.Errors:
// 			if !ok {
// 				return
// 			}
// 			logger.Error("WatchConfig", zap.Reflect("Error", err))
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }
