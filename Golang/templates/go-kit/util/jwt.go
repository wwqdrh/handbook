package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

/**
jwt-go 工具库的使用
 */

type UserClaim struct {
	Uname string `json:"username"`
	jwt.StandardClaims
}

func SampleGen() {
	sec := []byte("123abc")
	// 这里的是对称加密
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{Uname: "ahui"})
	token, _ := tokenObj.SignedString(sec)
	fmt.Println(token)

	uc := UserClaim{}
	getToken, _ := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
		return sec, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClaim).Uname)
		fmt.Println(getToken.Claims.(*UserClaim).ExpiresAt)
	}
}