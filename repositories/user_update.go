package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_gorm/model"
)

// Save会更新所有字段，即使你没有赋值
func UpdateSaveOrm(db *gorm.DB, user *model.User) {
	db.First(&user)

	user.Name = "oscar"
	user.Age = 30
	db.Save(&user)
}

// 更新指定字段，可以使用Update或者Updates
func UpdateFieldOrm(db *gorm.DB, user *model.User) {
	// 更新单个属性，如果它有变化
	// UPDATE user SET name='oscar', updated_at='2019-11-17 21:34:10' WHERE id=110;
	db.Model(&user).Update("name", "oscar")

	// 根据给定的条件更新单个属性
	// UPDATE users SET name='oscar', updated_at='2019-11-17 21:34:10' WHERE id=111 AND active=true;
	db.Model(&user).Where("active = ?", true).Update("name", "oscar")

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	// UPDATE users SET name='oscar', age=18, active=false, updated_at='2019-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Updates(map[string]interface{}{"name": "oscar", "age": 30, "active": false})

	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	// UPDATE users SET name='oscar', age=18, updated_at = '2019-11-17 21:34:10' WHERE id = 111;
	db.Model(&user).Updates(model.User{Name: "oscar", Age: 18})

	/*警告：当使用 struct 更新时，GORM只会更新那些非零值的字段*/
	// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
	db.Model(&user).Updates(model.User{Name: "", Age: 30})
}

// 更新选定字段	如果你想更新或忽略某些字段，你可以使用 Select，Omit
func UpdateSelectFieldOrm(db *gorm.DB, user *model.User) {
	// UPDATE users SET name='oscar', updated_at='2019-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "oscar", "age": 30})

	// UPDATE users SET age=18, name="oscar", updated_at='2019-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "oscar", "age": 30})
}

// 无 Hooks 更新
func UpdateHooksOrm(db *gorm.DB, user *model.User) {
	// 更新单个属性，类似于 `Update`
	// UPDATE users SET name='oscar' WHERE id = 111;
	db.Model(&user).UpdateColumn("name", "oscar")

	// 更新多个属性，类似于 `Updates`
	// UPDATE user SET name='oscar', age=18 WHERE id = 111;
	db.Model(&user).UpdateColumns(model.User{Name: "oscar", Age: 20})
}

// 批量更新 批量更新时 Hooks 不会运行
func UpdateBatchOrm(db *gorm.DB, user *model.User) {
	// UPDATE users SET name='oscar', age=18 WHERE id IN (10, 11);
	db.Table("user").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "oscar", "age": 20})

	/*使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}*/
	// UPDATE users SET name='oscar', age=30;
	db.Model(model.User{}).Updates(model.User{Name: "oscar", Age: 30})

	// 使用 `RowsAffected` 获取更新记录总数
	i := db.Model(model.User{}).Updates(model.User{Name: "oscar", Age: 20}).RowsAffected
	fmt.Println(i)
}

func UpdateSqlExpOrm(db *gorm.DB, order *model.Order) {
	// UPDATE "order" SET "price" = price * '2' + '100', "updated_at" = '2019-11-17 21:34:10' WHERE "id" = '2';
	db.Model(&order).Update("price", gorm.Expr("price * ? + ?", 2, 100))

	// UPDATE "order" SET "price" = price * '2' + '100', "updated_at" = '2019-11-17 21:34:10' WHERE "id" = '2';
	db.Model(&order).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})

	// UPDATE "order" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;
	db.Model(&order).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
}

func UpdateOtherOrm(db *gorm.DB, user *model.User) {
	// 为 update SQL 添加其它的 SQL
	db.Model(&user).Set("gorm:update_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Update("name", "oscar")
	// UPDATE users SET name='hello', updated_at = '2019-11-17 21:34:10' WHERE id=111 OPTION (OPTIMIZE FOR UNKNOWN);
}
