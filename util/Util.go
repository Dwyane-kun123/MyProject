/*
工具包
 */

package util

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
	"xukun/model"
)

func RandomString(n int) string  {
	letters := "qazwsxedcrfvtgbyhnujmiklopQAZWSXEDCRFVTGBYHNUJMIKOLP"
	res := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i,_ := range res{
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

func IsTelephineExit(db *gorm.DB,telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
