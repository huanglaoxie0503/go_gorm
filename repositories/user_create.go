package repositories

import (
	"github.com/jinzhu/gorm"
	"go_gorm/model"
)

func CreateOrm(db *gorm.DB) {
	user := &model.User{
		Model:        gorm.Model{},
		Name:         "",
		Age:          0,
		Birthday:     nil,
		Email:        "",
		Role:         "",
		MemberNumber: "",
		Num:          0,
		Address:      "",
		IgnoreMe:     0,
	}
	// 主键为空返回`true`
	db.NewRecord(user)

	db.Create(&user)

	// 创建`user`后返回`false`
	db.NewRecord(&user)
}
