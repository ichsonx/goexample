package main

import (
	"github.com/dgrijalva/jwt-go"
	"goexample/mycommon"
	"io/ioutil"
	"fmt"
	"time"
	"go/token"
)

const (
	//密钥用于签名加密
	privKeyPath = "./examplefiles/jwt.rsa"
	//公钥用于解密并验证签名
	pubKeyPath = "./examplefiles/jwt.rsa.pub"
)

var SignKey, VerifyKey []byte

func init() {
	var err error
	//读取密钥内容（二进制数组）
	SignKey, err = ioutil.ReadFile(privKeyPath)
	mycommon.Check(err)

	//读取公钥内容（二进制数组）
	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	mycommon.Check(err)
}

//定义一个自己的claims。必须继承（或者叫内嵌）jwt.StandardClaims自己的一个一般claims类型，这个类型结构遵循jwt的claims标准。
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

// Create the Claims，这是个实例化自定义claims的例子，下文未必用到，可做参考
var MyClaims = MyCustomClaims{
	"bar",
	jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	},
}

func main() {
	//genToken()
	validToken()
}

//获取token字符串
func genToken() string  {
	//用于测试token时间过期的，过期时间1秒，程序会在生成token后sleep3秒
	//MyClaims.ExpiresAt = time.Now().Add(time.Second * 1).Unix()
	MyClaims.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	MyClaims.Foo = "sonx"
	//还有一种方法获得token：jwt.New(签名方法)，这种方法获得的signer(签名器)，最后还需要给signer赋予claims：signer.Claims = claims（已实例化的claims）
	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims )
	token, err := signer.SignedString(SignKey)
	mycommon.Check(err)

	fmt.Printf("使用密钥加密签名后的token字符串：%s \n", token)
	return token
}

//从token字符串获取token、claims，验证token，打印claims里的内容
func validToken()  {
	//如果是要在请求中获取token，jwt有个request包，
	//里面有2个方法都可以获取token：jwt/requesst.ParseFromRequestWithClaims和jwt/requesst.ParseFromRequest
	tokenstring := genToken()
	time.Sleep(time.Second * 3)
	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		fmt.Printf("token 通过验证, Foo值： %s  \n", claims["foo"])
	}else {
		fmt.Printf("token 验证不通过： \n")
		fmt.Println(err)
	}

}
