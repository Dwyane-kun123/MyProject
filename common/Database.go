package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
	"xukun/model"
)

//初始化数据库
var DB *gorm.DB
func InitDB() *gorm.DB{
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.localhost")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))



	db, err := gorm.Open(driverName,args)
	if err != nil {
		panic("连接数据库错误，err" + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate(&model.User{})

	//给DB赋值 然后由下面的Getdb()返回 因为BD是全局变量
	DB = db
	return db
}

//返回全局变量BD
func Getdb() *gorm.DB {
	return DB
}