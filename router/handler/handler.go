package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-web-gin-example/common"
	"go-web-gin-example/common/JWT"
	"go-web-gin-example/common/redis"
	"go-web-gin-example/common/utils"
)

// request 的拦截器
func UserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("xxx-Token")
		token, err := jwt.ParseWithClaims(tokenString, &JWT.CustomToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT.SecretKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(200, &common.CommonResult{Res: -1, Mes: "token验证失败"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*JWT.CustomToken)
		if !ok {
			c.JSON(200, &common.CommonResult{Res: -1, Mes: "token验证失败"})
			c.Abort()
			return
		}
		//log
		utils.LogOutPutRequest(c, claims.Phone+" "+claims.UserName)
		// set token after xxx
		c.Set("Token_UserId", claims.UserId)
		c.Set("Token_Phone", claims.Phone)
		c.Set("Token_UserName", claims.UserName)
		// redis 限流
		redis.IsPeriodLimiting(c, claims.Phone)
	}
}
