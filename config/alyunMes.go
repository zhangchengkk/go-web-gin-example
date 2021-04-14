package config

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/labstack/gommon/log"
	"os"
)

const accessKeyId = ""
const accessSecret = ""

func initAliyunMes() {
	var err error
	MesClient, err = dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
