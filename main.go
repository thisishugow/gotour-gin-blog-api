package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"yuwang.idv/go-tour/blog-api/global"
	"yuwang.idv/go-tour/blog-api/internal/model"
	"yuwang.idv/go-tour/blog-api/internal/routers"
	"yuwang.idv/go-tour/blog-api/pkg/logger"
	"yuwang.idv/go-tour/blog-api/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatal("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatal("init.setupLogger err: %v", err)
	}
}

// @title Golang GIN API
// @version 1.0
// @description An expandable template API
// @termsOfService http://github.com/thisishugow
func main() {
	global.Logger.Debugf("%s: go-programming-tour-book/%s", "測試測試", "blog-api")

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	listenAddr := fmt.Sprintf("Start to listen on: %s\n", fmt.Sprintf("http://127.0.0.1%s", s.Addr))
	global.Logger.Info(listenAddr)
	global.Logger.Info(fmt.Sprintf("Swag on: %s\n", fmt.Sprintf("http://127.0.0.1%s/swagger/index.html", s.Addr)))
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Fatalf("main.s.ListenAndServe() err:%v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +
			global.AppSetting.LogFileName +
			global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
