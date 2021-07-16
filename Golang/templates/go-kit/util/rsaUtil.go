package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

/**
生成公私钥文件
 */
func GenRSAPubAndPri(bits int, filepath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	err = ioutil.WriteFile(filepath+"/private.pem", pem.EncodeToMemory(priBlock), 0644)
	if err != nil {
		return err
	}
	fmt.Println("====私钥文件创建成功=====")

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block {
		Type: "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = ioutil.WriteFile(filepath+"/public.pem", pem.EncodeToMemory(publicBlock), 0644)
	if err != nil {
		return err
	}
	fmt.Println("=====公钥文件创建成功======")
	return nil
}


func ReadRSAFile() {
	priKeyBytest, err := ioutil.ReadFile("./pem/private.pem")
	if err != nil {
		log.Fatal("私钥文件读取失败")
	}
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(priKeyBytest)
	if err != nil{
		log.Fatal("私钥文件不正确")
	}
	pubKeyBytes, err := ioutil.ReadFile("./pem/public.pem")
	if err != nil {
		log.Fatal("公钥文件读取失败")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Fatal("公钥文件不正确")
	}

	// 非对称加密
	user := UserClaim{Uname: "ahui"}
	user.ExpiresAt = time.Now().Add(time.Second * 5).Unix()
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodRS256, user)
	token, _ := tokenObj.SignedString(priKey)
	fmt.Println(token)

	i := 1
	for {// 使用循环测试当过期之后的效果
		uc := UserClaim{}
		getToken, _ := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
			return pubKey, nil
		})
		if getToken.Valid {
			fmt.Println(getToken.Claims.(*UserClaim).Uname)
			fmt.Println(getToken.Claims.(*UserClaim).ExpiresAt)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0{
				fmt.Println("token过期或未启用")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		} else {
			fmt.Println("无法解析此token, err")
		}
		i++
		fmt.Println(i)
	}
}