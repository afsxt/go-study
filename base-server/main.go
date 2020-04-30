package main

import (
	"base-server/models"
	"base-server/pkg/setting"
	"base-server/pkg/util"
	"base-server/routers"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//-----------------------------------------------------------------------------

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.TraceLevel)
	fileName := fmt.Sprintf("%s.%s",
		setting.AppSetting.LogSaveName, setting.AppSetting.LogFileExt)
	file, err := os.OpenFile(fileName,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Infoln("Failed to log to file, using default stderr")
	}
}

func init() {
	setting.Setup()
	initLogger()
	models.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
