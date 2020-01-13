package repositories

import (
	"github.com/jinzhu/gorm"
	"go_gorm/model"
)

/*
	警告 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。
*/

func DeleteOrm(db *gorm.DB, user model.User) {
	// 删除现有记录
	db.Delete(&user)

	// 为删除 SQL 添加额外的 SQL 操作
	// DELETE from emails where id=10 OPTION (OPTIMIZE FOR UNKNOWN);
	db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&user)
}

func DeleteBatchOrm(db *gorm.DB, user *model.User) {
	// DELETE from emails where email LIKE "%oscar%";
	db.Where("email LIKE ?", "%oscar%").Delete(&user)

	// DELETE from emails where email LIKE "%oscar%";
	db.Delete(model.User{}, "email LIKE ?", "%oscar%")
}

/*
如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！ 当调用 Delete 方法时， 记录不会真正的从数据库中被删除，
只会将DeletedAt 字段的值会被设置为当前时间
*/

// 软删除
func DeletionSoftOrm(db *gorm.DB, user model.User) {
	// UPDATE user SET deleted_at="2019-10-29 10:23" WHERE id = 111
	db.Delete(&user)

	// UPDATE user SET deleted_at="2019-10-29 10:23" WHERE age = 20;
	db.Where("age = ?", 20).Delete(&model.User{})

	// 查询记录时会忽略被软删除的记录
	// SELECT * FROM user WHERE age = 20 AND deleted_at IS NULL;
	db.Where("age = 20").Find(&user)

	// Unscoped 方法可以查询被软删除的记录
	// SELECT * FROM user WHERE age = 20;
	db.Unscoped().Where("age = 20").Find(&user)
}

// 物理删除  Unscoped 方法可以物理删除记录
// DELETE FROM orders WHERE id=10;
func DeleteOtherOrm(db gorm.DB, order model.Order) {
	db.Unscoped().Delete(&order)
}
