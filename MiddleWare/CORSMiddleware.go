package MiddleWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
解决浏览器同源问题
给header写入允许访问的域名或者是方法
 */

func CORSMiddleware() (gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		ctx.Writer.Header().Set("Access-Control-Allow-Max-Age", "86400") //设置缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*") //设置可以通过访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")//表示可以访问全部
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//判断请求的方法是否为option请求 是就返回200，否则继续请求
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		}else{
			ctx.Next()
		}
	}
}