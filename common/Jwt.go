package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"xukun/model"
)

/*
JSON Web Token
JWT的本质就是一个字符串，它是将用户信息保存到一个Json字符串中，然后进行编码后得到一个JWT token
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjgsImV4cCI6MTY1MzQ0NzcxNiwiaWF0IjoxNjUyODQyOTE2LCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.9Xy-yfw9w98ddsUBXOsBHJpgjkPzHpjjWikduNPP8k0
token使用的加密协议. 编码人员在jwt的claims实例化的信息. 前面两部份加上jwtkey哈希的一个值
*/

//定义JWT加密的密钥
var jwtkey = []byte("a_secrect_crect")

//定义token的声明
type Claims struct {
	UserId uint
	jwt.StandardClaims //又加入了jwt.StandardClaims这个结构体的方法
}
//登录成功后调用方法发放token
func ReleaseToken(user model.User) (string, error) {
	//定义token的截止时间
	expirationTime := time.Now().Add(7*24*time.Hour) //此时expirationTime也可以调用时间函数里的方法了
	//user的类型是model.User定义好的结构体 和用户信息发生关联 ！！！！！！！！！！！！！！！！！！！！！！！！！
	//结构体里套结构体
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//token结束时间
			ExpiresAt: expirationTime.Unix(),
			//token发放时间 NOW
			IssuedAt: time.Now().Unix(),
			//定义谁发放的token
			Issuer: "oceanlearn.tech",
			Subject: "user token",
		},
	}
	//上面都是token的准备工作
	//产生一个token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //jwt.SigningMethodHS256是签名算法
	//使用密钥产生一个加密的token
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error){
	claims := &Claims{}
	//用jwt.ParseWithClaims 将 tokenString解析
	//因为是用jetkey密钥去生成的token 也要用密钥去解析
	token, err:= jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{},error) {
		return jwtkey, nil
	})
	//&{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjksImV4cCI6MTY1MzQ2OTMxMCwiaWF0IjoxNjUyODY0NTEwLCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.3Ftu5gb3x0kRhJwPYht_PbN1a
	//b450R9Q_u73EViO0fY 0xc000005038 map[alg:HS256 typ:JWT] 0xc000086a80 3Ftu5gb3x0kRhJwPYht_PbN1ab450R9Q_u73EViO0fY true}
	//这就是解析成功的token
	//fmt.Println(token)
	return token, claims, err
}