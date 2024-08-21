package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"nftPlantform/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func ConnectDB() error {
	config := config.LoadConfig()
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.MySQL.DBUser,
		config.MySQL.DBPassword,
		config.MySQL.DBHost,
		config.MySQL.DBPort,
		config.MySQL.DBName)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	// Configure the connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Database connected successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
