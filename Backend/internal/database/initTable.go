package database

import (
	"database/sql"
	"nftPlantform/config"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// 表结构定义
var tables = []string{
	`CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		wallet_address VARCHAR(42) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_wallet_address (wallet_address)
	)`,
	`CREATE TABLE IF NOT EXISTS nfts (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		token_id VARCHAR(66) UNIQUE NOT NULL,
		contract_address VARCHAR(42) NOT NULL,
		owner_id BIGINT UNSIGNED NOT NULL,
		creator_id BIGINT UNSIGNED NOT NULL,
		metadata_uri VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (owner_id) REFERENCES users(id),
		FOREIGN KEY (creator_id) REFERENCES users(id),
		INDEX idx_token_id (token_id),
		INDEX idx_contract_address (contract_address)
	)`,
	`CREATE TABLE IF NOT EXISTS orders (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		seller_id BIGINT UNSIGNED NOT NULL,
		buyer_id BIGINT UNSIGNED,
		nft_id BIGINT UNSIGNED NOT NULL,
		price DECIMAL(20, 8) NOT NULL,
		status ENUM('OPEN', 'COMPLETED', 'CANCELLED') NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (seller_id) REFERENCES users(id),
		FOREIGN KEY (buyer_id) REFERENCES users(id),
		FOREIGN KEY (nft_id) REFERENCES nfts(id),
		INDEX idx_status (status)
	)`,
	`CREATE TABLE IF NOT EXISTS transactions (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		order_id BIGINT UNSIGNED NOT NULL,
		tx_hash VARCHAR(66) UNIQUE NOT NULL,
		amount DECIMAL(20, 8) NOT NULL,
		gas_fee DECIMAL(20, 8) NOT NULL,
		status ENUM('PENDING', 'COMPLETED', 'FAILED') NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (order_id) REFERENCES orders(id),
		INDEX idx_tx_hash (tx_hash)
	)`,
	`CREATE TABLE IF NOT EXISTS user_balances (
		user_id BIGINT UNSIGNED PRIMARY KEY,
		balance DECIMAL(20, 8) NOT NULL DEFAULT 0,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`,
}

func InitTable() {
	config := config.LoadConfig()
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
	config.MySQL.DBUser,
	config.MySQL.DBPassword,
	config.MySQL.DBHost,
	config.MySQL.DBPort,
	config.MySQL.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// 创建表
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}

	log.Println("Database initialization completed successfully")
}