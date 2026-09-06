package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;comment:รหัสผู้ใช้งาน (Primary Key)" json:"id"`
	Username     string    `gorm:"type:varchar(100);unique;not null;comment:ชื่อบัญชีผู้ใช้งาน (Unique)" json:"username"`
	PasswordHash string    `gorm:"type:text;not null;comment:รหัสผ่านที่ผ่านการเข้ารหัส (Bcrypt)" json:"password_hash"`
	CreatedAt    time.Time `gorm:"autoCreateTime;comment:วันที่และเวลาที่สร้างข้อมูล" json:"created_at"`
}