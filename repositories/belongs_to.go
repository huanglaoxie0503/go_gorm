package repositories

import (
	"github.com/jinzhu/gorm"
	"go_gorm/model"
)

type UserInfo struct {
	gorm.Model
	Name string
}

/*
 belongs to 会与另一个模型建立一对一关系，因此声明的每一个模型实例都会”属于”另一个模型实例。
*/

// `ProFile` 属于 `User`， 外键是`UserID`
type ProFile struct {
	gorm.Model
	UserID int
	User   UserInfo
	Name   string
}

/*
Foreign Key，若要定义属于关系的外键必须存在, 默认外键使用所有者的类型名称及其主键。

对于上述例子，定义一个属于 User 的模型，外键应该是 UserID。

GORM 提供了自定义外键的方法，例如：
*/

// 外键
type ProFIle1 struct {
	gorm.Model
	// 将 UserRefer 作为外键
	User      UserInfo `gorm:"foreignkey:UserRefer"`
	userRefer uint
}

/*
对于一个 belongs to 关系，GORM 通常使用所有者的主键作为外键的值，对于上面例子，外键的值是 User 的 ID。

当你关联一个 profile 到一个 user 时，GORM 将保存 user 的 ID 到 profile 的 UserID 字段。

你可以用 association_foreignkey 标签来更改它，例如：
*/

// 关联外键
type UserInfo1 struct {
	gorm.Model
	Refer string
	Name  string
}

type ProFIle2 struct {
	gorm.Model
	Name string
	// 将 Refer 作为关联外键
	User      UserInfo1 `gorm:"association_foreignkey:Refer"`
	UserRefer string
}

// Belongs To 的使用
// 你可以使用 Related 查找 belongs to 关系

func BelongsTo(db gorm.DB, user model.User, file ProFile) {
	db.Model(&user).Related(&file)
}
