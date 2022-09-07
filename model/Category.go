package model

/*
定义文章分类的结构体
 */

type Category struct {
	ID *uint `json:"id" gorm:"primary_key"` //主键
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateAt Time `json:"create_at" gorm:"type:timestamp"` //创建文章时间 当字段有CreateAt字段在插入一段记录的时候会自动填入时间
	UpdateAt Time `json:"update_at" gorm:"type:timestamp"` //更新文章时间
}

