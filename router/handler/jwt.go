package handler

import (
	"github.com/labstack/gommon/log"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type CustomToken struct {
	jwt.StandardClaims
	UserName   string
	Phone      string
	UserId     int
}

const SecretKey = "youKwtPrimateKey"
const ExpireTime = 3600 *  24// token有效期

/*
	struct => tokenString
*/
func CreateToken(Phone string, UserName string,UserId int) (string, error) {
	claims := CustomToken{
		UserName:   UserName,
		Phone:      Phone,
		UserId:     UserId,
	}
	claims.SetExpiredAt(time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	return signedToken, nil
}

/*
	tokenString => token
*/
func VerfyToken(StrToken string) (*CustomToken, error) {
	tokenInfo := &CustomToken{}
	token, err := jwt.ParseWithClaims(StrToken, tokenInfo, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Error(err.Error())

		return nil, err
	}
	if err = token.Claims.Valid(); err != nil {
		return nil, err
	}
	return tokenInfo, nil
}



func (c *CustomToken) SetExpiredAt(expiredAt int64) {
	c.ExpiresAt = expiredAt
}