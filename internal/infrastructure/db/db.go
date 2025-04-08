package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewGormSQLServer() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	// DSN for the master database (needed to create target DB)
	masterDSN := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=master", user, password, host, port)

	masterDB, err := gorm.Open(sqlserver.Open(masterDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Could not connect to SQL Server master DB: %v", err)
	}

	// Ensure target database exists
	createDBSQL := fmt.Sprintf("IF DB_ID('%s') IS NULL CREATE DATABASE [%s];", database, database)
	if err := masterDB.Exec(createDBSQL).Error; err != nil {
		log.Fatalf("❌ Could not create database %s: %v", database, err)
	}

	// Connect to the target database
	targetDSN := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, database)

	db, err := gorm.Open(sqlserver.Open(targetDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Could not connect to database %s: %v", database, err)
	}

	log.Println("✅ Connected to SQL Server:", database)
	return db
}
