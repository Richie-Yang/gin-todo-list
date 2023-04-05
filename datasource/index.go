package datasource

import (
	"fmt"
	"gin-todo-list/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	ip := os.Getenv("DB_IP")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, name)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	config := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(config), &gorm.Config{})

	if err == nil {
		DB = db
	} else {
		println("%v", err)
		os.Exit(1)
	}
}

func Migrate() {
	Setup()
	DB.AutoMigrate(&models.TodoList{}, &models.Todo{})
}
