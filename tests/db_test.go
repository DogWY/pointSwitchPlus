package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name string
	Pwd  string
}

func TestDB(t *testing.T) {
	//utils.DB.AutoMigrate(models.User{})
	db, err := gorm.Open(mysql.Open("root:123456@tcp(8.130.50.95:3306)/project?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("connect fail")
	}
	db.AutoMigrate(User{})
	user := User{
		Name: "tqh",
		Pwd:  "123456",
	}
	db.Model(&user).Create(&user)
}
