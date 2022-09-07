package MiddleWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xukun/common"
	"xukun/model"
)

//验证token的中间件
func AuthMiddleWare() (gin.HandlerFunc) { //路由的包装处理函数
	return func(ctx *gin.Context) {
		//获取Authorization Header 获取头部请求
		tokenString:= ctx.GetHeader("Authorization")
		//验证token格式
		//如果tokestring为空或者头部不是bearer
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){ //这应该是客户端定义的方式，服务端是来验证的
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401,"msg":"权限不足"})
			//程序终止
			ctx.Abort()
			return
		}

		//返回正确的token
		tokenString = tokenString[7:]
		//要对token解密
		token, claimss, err := common.ParseToken(tokenString) //claims是对结构体Claim的实例化
		//解析token成功就获取用户信息 如果返回失败 或者 返回错误 就关闭请求
		if err != nil || !token.Valid{
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		//token通过验证，获取claims中的userID
		userId := claimss.UserId //uint类型
		DB:=common.Getdb()
		var user model.User
		//找到符合userId这个条件的
		DB.First(&user,userId)

		//用户
		if user.ID == 0{ //ID        uint `gorm:"primary_key"`  这是主键
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401,"msg":"权限不足"})
			ctx.Abort()
			return
		}

		//用户存在 将user的信息写入上下文 这样GET函数才能取到
		ctx.Set("user",user)
		//先执行下文  再返回来执行本函数里接下来的语句
		ctx.Next()
	}
}