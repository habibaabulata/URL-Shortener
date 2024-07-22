package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string `gorm:"type:varchar(100);uniqueIndex"`
    Password string
}
