package models

import "go-web-gin-example/config"

type User struct {
	Id         int     `xorm:"int autoincr pk"`
	Phone      string `xorm:"varchar(20) notnull unique default '' "`
	UserName   string `xorm:"varchar(20) notnull default '' "`
	Password   string `xorm:"varchar(20) notnull default '123456' "`
}


func QueryUserByPhone(phone string) (*User, error) {
	user := &User{}
	_, err := config.Engine.
		Where("phone = ?").
		Cols("user_name", "password", "phone", "id").
		Get(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

