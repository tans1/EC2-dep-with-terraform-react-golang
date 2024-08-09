package config

import (
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
)

type IDbConfig interface {
	CreateDB()
	ConnectDb()
}

type DbConfig struct{}

func NewDbConfig() *DbConfig {
	return &DbConfig{}
}

func (c *DbConfig) CreateDB() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)
	db, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		log.Print("failed to connect to DB server",err)
		return err
	}

	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := db.Exec(createDBQuery).Error; err != nil {
		log.Print("failed to create database",err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Print("failed to get DB from gorm",err)
		return err
	}
	sqlDB.Close()
	return nil

}

func (c *DbConfig) ConnectDb() (*gorm.DB, error) {
	if err := c.CreateDB(); err != nil {
		return nil, err
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Print("failed to connect database schema")
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Blog{}, &domain.Comment{}, &domain.Reply{})
	if err != nil {
		log.Print("failed to migrate database schema")
		return nil, err
	}
	return db, err
}
