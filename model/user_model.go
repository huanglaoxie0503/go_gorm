package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type LikeNum struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

type User struct {
	gorm.Model
	Name string
	Age int64
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"`
	// 设置字段大小为255
	Role string `gorm:"size:225"`
	// 设置会员号（member number）唯一并且不为空
	MemberNumber string `gorm:"unique;not null"`
	// 设置 num 为自增类型
	Num int `gorm:"AUTO_INCREMENT"`
	// 给address字段创建名为addr的索引
	Address string `gorm:"index:addr"`
	// 忽略本字段
	IgnoreMe int `gorm:"-"`
}

type Model struct {
	ID uint `gorm:"primary_key"`

}

type Order struct {

}