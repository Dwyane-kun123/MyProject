package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"xukun/common"
)

/*
1、在执行登录成功的时候，会生成一个token（用于接下来的身份校验）
2、在进行接下来的通讯的时候，客户端不需要再继续连接，只需要向服务端发送token就可以了
3、验证token的流程是这样的： ① 服务端在登陆成功后用密钥生成token ② 客户端在请求信息的时候进行token验证
   ③ 先把返回生成的token，接着服务端用密钥解密，如果验证成功就调用数据库里的相关信息，这是用ID作为主键生成的
 */

func main() {
	InitConfig() //在项目启动的一开始就读取我们的配置
	db := common.InitDB()
	defer db.Close()

	r := gin.Default() //
	//在controller包里封装好了Register
	r = ControllerRouter(r)


	//配置一下端口号
	port := viper.GetString("server.port")
	if  port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig()  {
	//获取工作目录
	workDir,_ := os.Getwd()
	//设置要读取的文件名
	viper.SetConfigName("application")
	//设置要读取的文件的类型
	viper.SetConfigType("yml")
	//设置文件的路径
	viper.AddConfigPath(workDir + "/config")
	err:= viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

