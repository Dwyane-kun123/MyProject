package cotroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"xukun/common"
	"xukun/model"
	"xukun/response"
)
/*
根据文章的ID进行分类
 */

//对数据库作操作 因为要CRUD
type CategoryController struct {
	DB *gorm.DB
}

//用接口来封装方法
type ICategoryController interface {
	RestController //复用RestController里的代码 以便于后续的操作
}

//用这个将CategoryController实例化 返回的是一个接口 在路由中用这个实例化来做调用
//为了和数据库交互
func NewCategoryController() ICategoryController {
	db := common.Getdb()
	//自动迁移
	db.AutoMigrate(model.Category{})
	//用在数据库里取到的db实例化CategoryController
	return CategoryController{DB:db}
}

//func (c CategoryController) Create(ctx *gin.Context)  {
//	//用结构体接受数据
//	//var requestCategory model.Category
//	//ctx.Bind(requestCategory)
//	nameTxt := ctx.PostForm("name")
//
//	log.Println(nameTxt)
//	if nameTxt == "" {
//		response.Fail(ctx,nil, "数据验证错误，分类名称必填")
//		return
//	}
//
//	requestCategory := model.Category{
//		Name: nameTxt,
//		CreateAt: time.Now(),
//		UpdateAt: time.Now(),
//	}
//	c.DB.Create(&requestCategory)
//
//	response.Success(ctx, gin.H{"category":requestCategory},"")
//}

func (c CategoryController) Create(ctx *gin.Context)  {
	//用结构体接受数据
	var requestCategory model.Category
	ctx.Bind(&requestCategory)


	log.Println(requestCategory.Name)
	if requestCategory.Name == "" {
		response.Fail(ctx,nil, "数据验证错误，分类名称必填")
		return
	}

	//自动生成数据库的category表
	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"category":requestCategory},"")
}

func (c CategoryController) Update(ctx *gin.Context)  {
	//绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)


	log.Println(requestCategory.Name)
	if requestCategory.Name == "" {
		response.Fail(ctx,nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path中的参数
	categoryid,_ := strconv.Atoi(ctx.Params.ByName("id"))

	//先查找一下有没有要修改的值
	var updateCategory model.Category
	if c.DB.First(&updateCategory,categoryid).RecordNotFound(){
		response.Fail(ctx,nil,"分类不存在")
	}

	//更新分类
	//从前端传过来的值可以是map,struct，name value
	                                     //需要修改的字段    需要修改的值
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory},"修改成功")
}

func (c CategoryController) Show(ctx *gin.Context)  {
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	//看看能不能在数据库中找到同ID的数据
	var Category model.Category
	if c.DB.First(&Category,categoryId).RecordNotFound(){
		response.Fail(ctx,nil,"分类不存在")
	}

	response.Success(ctx, gin.H{"category": Category},"查看成功")
}

func (c CategoryController) Delete(ctx *gin.Context)  {
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	//如果删除失败
	if err := c.DB.Delete(model.Category{},categoryId).Error; err!= nil{
		response.Fail(ctx,nil,"删除不成功")
		return
	}

	response.Success(ctx,nil,"删除成功")
}