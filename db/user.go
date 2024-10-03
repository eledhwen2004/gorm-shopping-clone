package db

import (
	"time"

	"gorm.io/gorm"
)

type userRole int

const (
	admin userRole = iota
	supplier
	customer
)

type User struct {
	gorm.Model
	ID          string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Username    string         `json:"username" gorm:"not null"`
	Password    string         `json:"password" gorm:"not null"`
	Email       string         `json:"email" gorm:"not null"`
	PhoneNumber string         `json:"phone_number" gorm:"not null"`
	Role        userRole       `json:"role" gorm:"not null"`
	Warnings    uint           `json:"warning" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Customer struct {
	gorm.Model
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	UserID    string    `json:"user_id" gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;not null"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Comments  []Comment `gorm:"foreignKey:CommentID"`
}

type Supplier struct {
	gorm.Model
	ID          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	UserID      string    `json:"user_id" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;not null"`
	CompanyName string    `json:"company_name" gorm:"not null"`
	Address     string    `json:"address" gorm:"not null"`
	Products    []Product `gorm:"many2many:product_suppliers"`
}

func createAccount() {
}

func deleteAccount() {}

func updateAccount() {}

func getAccount() {}
