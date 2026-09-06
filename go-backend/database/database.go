package database

import (
	"log"
	"time"

	"example.com/interview-question-087/go-backend/config"
	"example.com/interview-question-087/go-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := config.AppConfig.GetDSN()

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if !db.Migrator().HasTable(&models.User{}) {
		log.Println("Table 'users' does not exist. Creating table with Auto Increment ID...")
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("Failed to create 'users' table: %v", err)
		}
		db.Exec("COMMENT ON TABLE users IS 'ตารางเก็บข้อมูลผู้ใช้งานระบบ';")
		log.Println("Table 'users' created successfully!")
	} else {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Printf("Auto migration warning: %v", err)
		}
		log.Println("Database migration verified: 'users' table is ready")
	}

	DB = db
	log.Println("Database connection established successfully")
}