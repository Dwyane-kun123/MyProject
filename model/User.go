package model

import "github.com/jinzhu/gorm"

type User struct {
	//gorm默认以ID为主键
	gorm.Model        //代码中定义模型（Models）与数据库中的数据表进行映射
	Name       string `gorm:"type:varchar(20);not null" json:"name"`
	Telephone  string `gorm:"type:varchar(11);not null" json:"telephone"`
	Password   string `gorm:"size:255";not null json:"password"`
}
