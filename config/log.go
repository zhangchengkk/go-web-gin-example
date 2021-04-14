package config

import (
	"fmt"
	"io"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func initLog() {
	isExisted := isExist(LogPath)
	if isExisted == false {
		err := os.Mkdir(LogPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	logPath := LogPath + "/gin_" + time.Now().Format("2006-01-02 15:04:05") + ".log"
	// Logging to a file.
	f, err := os.Create(logPath)
	if err != nil {
		log.Error(err)
	}
	if Env == "pro" {
		gin.DefaultWriter = io.MultiWriter(f)
	}
	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

