package model

import (
	"fmt"
	"server/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

func Login(username string, password string) (msg string, token string) {
	type User struct {
		id       int
		Username string
		Password string
	}

	db := utils.DB
	user := User{}

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			msg = "账号错误"
		} else {
			msg = "登录过程中出现错误..."
		}
		return
	}

	if user.Password != password {
		msg = "密码错误"
	} else {
		msg = "登录成功"
		token = CreateJWT()
	}

	return
}

func CreateJWT() (jwtstr string) {
	c := jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 60,   // 生效时间,当前时间的60s之前
		ExpiresAt: time.Now().Unix() + 3600, // 失效时间,当前时间的1个小时后
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	mySigningKey := []byte("iamlulu")
	jwtstr, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	return
}
