package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_USERNAME = "postgres"
var DB_PASSWORD = "admin"
var DB_NAME = "ta_micro_pendataan"
var DB_HOST = "127.0.0.1"
var DB_PORT = "5432"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error

	dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable TimeZone=Asia/Jakarta"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("Error connecting to database : error=", err)
		return nil
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return db
}
