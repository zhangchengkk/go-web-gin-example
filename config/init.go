package config

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
)
var MesClient *dysmsapi.Client
var Engine *xorm.Engine
var RedisClient *redis.Client

func init() {
	initLog()
	initDb()
	//initRedis()
	//initAliyunMes()
	//initAliPay()
}