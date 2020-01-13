package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_gorm/model"
	"time"
)

// 查询
func QueryOrm(db *gorm.DB, user *model.User) {
	// 根据主键查询第一条记录 -->	SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user)
	// 随机获取一条记录 -->	SELECT * FROM users LIMIT 1;
	db.Take(&user)
	// 根据主键查询最后一条记录 -->	SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	// 查询所有的记录 --> 	SELECT * FROM users;
	db.Last(&user)
	// 查询指定的某条记录(仅当主键为整型时可用)  -->	SELECT * FROM users WHERE id = 10;
	db.First(&user, 10)
}

// Where 条件查询
func QueryWhereOrm(db *gorm.DB, user *model.User) {
	// SELECT * FROM users WHERE name = 'oscar' limit 1;
	db.Where("name = ?", "oscar").First(&user)

	// SELECT * FROM users WHERE name = 'oscar';
	db.Where("name = ?", "oscar").Find(&user)

	// SELECT * FROM users WHERE name <> 'oscar';
	db.Where("name <> ?", "oscar").Find(&user)

	// SELECT * FROM users WHERE name in ('oscar','oscar_01 2');
	db.Where("name IN (?)", []string{"oscar", "oscar_01"}).Find(&user)

	// SELECT * FROM users WHERE name LIKE '%os%';
	db.Where("name LIKE ?", "%os%").Find(&user)

	// SELECT * FROM users WHERE name = 'oscar' AND age >= 22;
	db.Where("name = ? AND age >= ?", "oscar", "25").Find(&user)

	// SELECT * FROM users WHERE updated_at > '2020-01-10 00:00:00';
	friday := time.Friday
	db.Where("updated_at > ?", friday).Find(&user)

	// between...and...
	db.Where("age between ? and ?", 10, 30).Find(&user)
}

// 当通过结构体进行查询时，GORM将会只通过非零值字段查询，这意味着如果你的字段值为0，''， false 或者其他 零值时，将不会被用于构建查询条件
func QueryStructMapOrm(db *gorm.DB, user *model.User) {
	// Struct	SELECT * FROM users WHERE name = "oscar" AND age = 30 LIMIT 1;
	db.Where(&model.User{Name: "oscar", Age: 30}).First(&user)

	// Map	SELECT * FROM users WHERE name = "oscar" AND age = 30;
	db.Where(map[string]interface{}{"name": "oscar", "age": 30}).Find(&user)

	// 主键的切片	SELECT * FROM users WHERE id IN (20, 21, 22);
	db.Where([]int64{20, 21, 22}).Find(&user)
}

// Not 条件
func QueryNotOrm(db *gorm.DB, user *model.User) {
	// SELECT * FROM users WHERE name <> "oscar" LIMIT 1;
	db.Not("name", "oscar").First(&user)

	// SELECT * FROM users WHERE id NOT IN (1,2,3);	Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)

	// SELECT * FROM users;
	db.Not([]int64{}).First(&user)

	// SELECT * FROM users WHERE NOT(name = "oscar");
	db.Not("name = ?", "oscar").First(&user)

	// SELECT * FROM users WHERE name <> "oscar";
	db.Not(model.User{Name: "oscar"}).First(&user)
}

// Or 条件
func QueryOrOrm(db *gorm.DB, user *model.User) {
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&user)

	// Struct
	db.Where("name = ?").Or(model.User{Name: "oscar_2"}).Find(&user)

	// Map	SELECT * FROM users WHERE name = 'oscar' OR name = 'oscar_2';
	db.Where("name = 'oscar'").Or(map[string]interface{}{"name": "oscar_2"}).Find(&user)
}

// Inline Condition 内联条件
func QueryConditionOrm(db *gorm.DB, user *model.User) {
	// 根据主键获取记录 (只适用于整形主键)	SELECT * FROM users WHERE id = 23 LIMIT 1;
	db.First(&user, 23)

	// 根据主键获取记录, 如果它是一个非整形主键	SELECT * FROM users WHERE id = '600519' LIMIT 1;
	db.First(&user, "id = ?", "600519")

	// SELECT * FROM users WHERE name = "oscar";
	db.Find(&user, "name = ?", "oscar")

	// Struct SELECT * FROM users WHERE age = 30;
	db.Find(&user, model.User{Age: 30})

	// Map  SELECT * FROM users WHERE age = 30;
	db.Find(&user, map[string]interface{}{"age": 30})
}

// 其它查询选项
func QueryExtraQueryingOption(db *gorm.DB, user *model.User) {
	// 为查询 SQL 添加额外的 SQL 操作	SELECT * FROM users WHERE id = 10 FOR UPDATE;
	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
}

// 获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象 (仅支持 struct 和 map 条件)
func QueryFirstOrInit(db *gorm.DB, user *model.User) {
	// 未找到	user -> User{Name: "non_existing"}
	db.FirstOrInit(&user, model.User{Name: "non existing"})

	// 找到	user -> User{Id: 1, Name: "oscar", Age: 30}
	db.Where(model.User{Name: "oscar"}).FirstOrInit(&user)

	// user -> User{Id: 111, Name: "oscar", Age: 30}
	db.FirstOrInit(&user, map[string]interface{}{"name": "oscar"})

}

// 如果记录未找到，将使用参数初始化 struct
func QueryAttrsOrm(db *gorm.DB, user *model.User) {
	// 未找到	SELECT * FROM USERS WHERE name = 'non_existing';	user -> User{Name: "non_existing", Age: 30}
	db.Where(model.User{Name: "non existing"}).Attrs(model.User{Age: 30}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = 'non_existing';		user -> User{Name: "non_existing", Age: 30}
	db.Where(model.User{Name: "non existing"}).Attrs("age", 30).FirstOrInit(&user)

	// 找到	SELECT * FROM USERS WHERE name = oscar';	user -> User{Id: 111, Name: "oscar", Age: 30}
	db.Where(model.User{Name: "oscar"}).Attrs(model.User{Age: 30}).FirstOrInit(&user)
}

// 不管记录是否找到，都将参数赋值给 struct
func QueryAssignOrm(db *gorm.DB, user *model.User) {
	// 未找到 	user -> User{Name: "non_existing", Age: 20}
	db.Where(model.User{Name: "non_existing"}).Assign(model.User{Age: 30}).FirstOrInit(&user)

	// 找到	SELECT * FROM USERS WHERE name = oscar';	user -> User{Id: 111, Name: "oscar", Age: 20}
	db.Where(model.User{Name: "oscar"}).Assign(model.User{Age: 30}).FirstOrInit(&user)
}

// 获取匹配的第一条记录, 否则根据给定的条件创建一个新的记录 (仅支持 struct 和 map 条件)
func QueryFirstOrCreateOrm(db *gorm.DB, user *model.User) {
	// 未找到	INSERT INTO "users" (name) VALUES ("non_existing");	user -> User{Id: 112, Name: "non_existing"}
	db.FirstOrCreate(&user, model.User{Name: "non_existing"})

	// 找到	user -> User{Id: 111, Name: "user"}
	db.Where(model.User{Name: "oscar"}).FirstOrCreate(&user)

	// 1. Attrs 如果记录未找到，将使用参数创建 struct 和记录
	// (1). 未找到
	// SELECT * FROM users WHERE name = 'non_existing';
	// INSERT INTO "users" (name, age) VALUES ("non_existing", 30);
	// user -> User{Id: 1, Name: "non_existing", Age: 30}
	db.Where(model.User{Name: "non_existing"}).Attrs(model.User{Age: 30}).FirstOrCreate(&user)

	// (2). 找到
	// SELECT * FROM users WHERE name = 'oscar';
	// user -> User{Id: 111, Name: "oscar", Age: 20}
	db.Where(model.User{Name: "oscar"}).Attrs(model.User{Age: 30}).FirstOrCreate(&user)

	// 2. Assign 不管记录是否找到，都将参数赋值给 struct 并保存至数据库.
	// (1).未找到
	// SELECT * FROM users WHERE name = 'non_existing';
	// INSERT INTO "users" (name, age) VALUES ("non_existing", 30);
	// user -> User{Id: 112, Name: "non_existing", Age: 30}
	db.Where(model.User{Name: "non_existing"}).Assign(model.User{Age: 30}).FirstOrCreate(&user)

	// (2). 找到
	// SELECT * FROM users WHERE name = 'oscar';
	// UPDATE users SET age=30 WHERE id = 111;
	// user -> User{Id: 111, Name: "oscar", Age: 30}
	db.Where(model.User{Name: "oscar"}).Assign(model.User{Age: 30}).FirstOrCreate(&user)
}

// 高级查询
func QueryAdvancedOrm(db *gorm.DB, order *model.Order) {
	// 1. SubQuery 子查询  基于 *gorm.expr 的子查询
	// SELECT * FROM "order"  WHERE "order"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "order"  WHERE (state = 'paid')));
	db.Where("amount > ?", db.Table("order").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&order)

	// 2. Select 指定你想从数据库中检索出的字段，默认会选择全部字段。
	// SELECT name, price FROM order;
	db.Select("name", "price").Find(&order)

	// SELECT name, price FROM order
	db.Select([]string{"name", "price"}).Find(&order)

	// SELECT COALESCE(price,'100') FROM users;
	_, _ = db.Table("order").Select("COALESCE(price,?)", 100).Rows()

	// 3. Order 指定从数据库中检索出记录的顺序。设置第二个参数 reorder 为 true ，可以覆盖前面定义的排序条件。
	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("price desc, name").Find(&order)

	// 多字段排序
	db.Order("price desc").Order("name").Find(&order)

	// 覆盖排序
	// SELECT * FROM users ORDER BY age desc; (order1)
	// SELECT * FROM users ORDER BY age; (order2)
	order2 := model.Order{}
	db.Order("price desc").Find(&order).Order("price", true).Find(&order2)

	// 4. Limit 数量
	// SELECT * FROM users LIMIT 3;
	db.Limit(3).Find(&order)

	// -1 取消 Limit 条件
	// SELECT * FROM order LIMIT 10; (order)
	// SELECT * FROM order; (order2)
	db.Limit(10).Find(&order).Limit(-1).Find(&order2)

	// 5. Offset 偏移 指定开始返回记录前要跳过的记录数
	// SELECT * FROM users OFFSET 3;
	db.Offset(3).Find(&order)
	// -1 取消 Offset 条件
	// SELECT * FROM order OFFSET 10; (order)
	// SELECT * FROM order; (order2)
	db.Offset(10).Find(&order).Offset(-1).Find(&order2)

	// 6.Count 总数 该 model 能获取的记录总数。
	var count *int
	// SELECT * from USERS WHERE name = 'oscar' OR name = 'oscar 2'; (order)
	// SELECT count(*) FROM users WHERE name = 'oscar' OR name = 'oscar 2'; (count)
	db.Where("name = ?", "oscar").Or("name = ?", "oscar_02").Find(&order).Count(&count)

	// SELECT count(*) FROM order WHERE name = 'oscar'; (count)
	db.Model(&model.User{}).Where("name = ?", "oscar").Count(&count)

	// SELECT count(*) FROM deleted_users;
	db.Table("deleted_users").Count(&count)

	// SELECT count( distinct(name) ) FROM deleted_users;
	// Count 必须是链式查询的最后一个操作 ，因为它会覆盖前面的 SELECT，但如果里面使用了 count 时不会覆盖
	db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
}

// Group & Having
func QueryGroupHaving(db *gorm.DB, order *model.Order) {
	rows1, err1 := db.Table("order").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
	if err1 != nil {
		fmt.Println(err1)
	}
	for rows1.Next() {

	}

	rows2, err2 := db.Table("order").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	if err2 != nil {
		fmt.Println(err2)
	}
	for rows2.Next() {

	}

	db.Table("order").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&order)
}

func QueryJoinsOrm(db *gorm.DB, order *model.Order) {
	rows, err := db.Table("user").Select("user.name, email.email").Joins("left join email.user_id = user.id").Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {

	}

	db.Table("user").Select("user.name, email.email").Joins("left join email on email.user_id = user.id").Scan(&order)

	db.Joins("JOIN emails ON emails.user_id = user.id AND emails.email = ?", "oscar@qq.com").Joins("JOIN credit_cards ON credit_cards.user_id = user.id").Where("credit_cards.number = ?", 110).Find(&order)
}

// Pluck，查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
func QueryPluckOrm(db *gorm.DB, user *model.User) {
	var ages []int64
	db.Find(&user).Pluck("age", &ages)

	var names []string
	db.Model(&model.User{}).Pluck("name", &names)

	db.Table("deleted_users").Pluck("name", &names)

	// 查询多个字段
	db.Select("name, age").Find(&user)
}

// Scan，扫描结果至一个 struct
func QueryScanOrm(db *gorm.DB) {
	type Result struct {
		Name string
		Age  int
	}
	var result Result
	db.Table("user").Select("name, age").Where("name = ?", "oscar").Scan(&result)

	db.Raw("select name, age from user where name = ?", "oscar").Scan(&result)
}
