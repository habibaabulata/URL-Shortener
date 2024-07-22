package database

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "url-shortener/config"
    "url-shortener/models"
)

var DB *gorm.DB

func InitDB() {
    dsn := config.GetDSN()
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v\n", err)
    }

    DB.AutoMigrate(&models.User{}, &models.URL{})
}
