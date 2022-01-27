package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	db_user := os.Getenv("DB_USER")
	// db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_host, db_port, db_name)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}

	db = conn

}

func GetDB() *gorm.DB {
	return db
}
