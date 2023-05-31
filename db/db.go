package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return nil
	}
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DEV_USERNAME") + ":" + os.Getenv("DB_DEV_PASSWORD") + "@tcp" + "(" + os.Getenv("DB_DEV_HOST") + ":" + os.Getenv("DB_DEV_PORT") + ")/" + os.Getenv("DB_DEV_NAME") + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Error connecting to database : error=", err)
		return nil
	}

	return db
}
