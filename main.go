package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_gorm/repositories"
)

func main() {
	db, err := repositories.Connection()
	if err != nil {
		panic(err)
	}
	//repositories.CreatTableOrm(db)
	repositories.InsertOrm(db)
}
