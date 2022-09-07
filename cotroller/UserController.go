package cotroller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"xukun/common"
	"xukun/dto"
	"xukun/model"
	"xukun/response"
	"xukun/util"
)


//注册功能
func Register(ctx *gin.Context) {
	db := common.Getdb()

	//用MAP获取请求的参数
	//RequestMap := make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&RequestMap)

	//用结构体来接受参数的请求
	var RequestUser model.User
	//json.NewDecoder(ctx.Request.Body).Decode(&RequestUser)
	//gin框架提供的bind函数来接受从前端返回的json数据
	ctx.Bind(&RequestUser)
	//
	//name:= ctx.PostForm("name") //这是直接从
	//telephone:= ctx.PostForm("telephone")
	//password:= ctx.PostForm("password")
	name :=  RequestUser.Name
	telephone := RequestUser.Telephone
	password := RequestUser.Password

	log.Println(name,telephone)
	if len(telephone) != 11{
		//fmt.Println(len(telephone),telephone)
		//fmt.Println(password)
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"电话必须是11位"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "电话必须是11位")
		return
	}
	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"密码不能小于6位"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if len(name) == 0  {
		name = util.RandomString(10)
	}

	//查看手机号是否存在
	if util.IsTelephineExit(db, telephone){ //这是在查询数据库了
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"用户已经存在，不允许注册"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在，不允许注册")
		return
	}
	//创建用户
	//将密码加密存储
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		//返回系统级的错误
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":500,"msg":"加密错误"})
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		//操作直接结束了
		return
	}
	//User结构体是调用了gorm.model结构体里的字段, 上述条件都满足的话 POST到的数据存入数据库
	newUser:= model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasPassword),
	}

	//根据字段名创建新的用户
	db.Create(&newUser)
	//返回结果
	//ctx.JSON(200,gin.H{
	//	"msg":"注册成功",
	//	"code": 200,
	//})

	//用登录信息直接返回token
	token,errr:= common.ReleaseToken(newUser)
	if errr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code":500,"msg":"系统异常"})
		//记录日志，将错误的信息打印出来
		log.Printf("token generate error: %v",errr)
		return
	}


	//response.Success(ctx,nil,"注册成功")

	response.Success(ctx, gin.H{"token": token}, "注册成功")
	log.Println(name,telephone,password)
	return
}

//登录功能
func Login(ctx *gin.Context)  {

	//创建一个函数连接数据库
	DB := common.Getdb()
	//获取参数
	telephone:= ctx.PostForm("telephone")
	password:= ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11{
		//fmt.Println(len(telephone),telephone)
		//fmt.Println(password)
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"电话必须是11位"})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "电话必须是11位")
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"密码不能小于6位"})
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)  //将数据库里的主键取出来做判断 数据库中具体的数据是根据telephone取出来的
	//
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422,"msg":"用户不存在"})
		return
	}
	//判断密码是否正确，比较原始的密码和加密后的密码是否一样
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err!= nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code":400,"msg":"密码验证错误"})
		return
	}

	//发放token

	token,errr:= common.ReleaseToken(user)
	if errr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code":500,"msg":"系统异常"})
		//记录日志，将错误的信息打印出来
		log.Printf("token generate error: %v",errr)
		return
	}

	//如果登录成功，返回结果
	//ctx.JSON(200,gin.H{
	//	"code": 200,
	//	"msg":"登录成功",
	//	"token" : token,
	//})
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

//从上下文中获取用户信息
func Info(ctx *gin.Context)  {
	//从上下文获取用户信息，然后返回用户信息
	user,_ := ctx.Get("user") //ctx.Get("user") 跨中间件取值

	ctx.JSON(http.StatusOK, gin.H{"code":200, "data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
	//user.(model.User) 有点像断言的意思，只找出想要的返回
}




