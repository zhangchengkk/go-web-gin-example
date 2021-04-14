package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-gin-example/common"
	"net/http"
)

func LoginApi(c *gin.Context) {
	//passwd := c.Query("passwory")
	//get from token
	//phone, _ := c.Get("Token_Phone")
	//_, _ = fmt.Fprintln(gin.DefaultWriter, "login request: " + phone.(string) + passwd)
	c.JSON(http.StatusOK, &common.CommonResult{Res: 0, Mes: "ok", Data: "custom interface{} data"})
}

func RefreshUserInfoApi(c *gin.Context) {
	userId := c.Query("userId")
	_, _ = fmt.Fprintln(gin.DefaultWriter, "refresh info" + userId)
	c.JSON(http.StatusOK, &common.CommonResult{Res: 0, Mes: "ok", Data: "custom interface{} data"})
}
