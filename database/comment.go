package database

import (
	"fmt"
	"time"

	"shopping-clone/postgre"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID         string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Content    string         `json:"content" gorm:"type:varchar(100)"`
	CustomerID string         `json:"customer_id" gorm:"not null"`
	Customer   Customer       `gorm:"foreignKey:CustomerID;not null"`
	ProductID  string         `json:"product_id" gorm:"not null"`
	Product    Product        `gorm:"foreignKey:ProductID;not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func CreateComment(comment *Comment) error {
	result := postgre.DB.Create(&comment)
	if result.Error != nil {
		fmt.Printf("Error while creating comment %v\n", result.Error)
	}
	return result.Error
}

func ReadComment(commentID string) (*Comment, error) {
	comment := Comment{}
	result := postgre.DB.Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		fmt.Printf("Error while finding comment %v\n", result.Error)
	}
	return &comment, result.Error
}

func UpdateComment(comment *Comment) error {
	result := postgre.DB.Model(&Comment{}).Where("id = ?", comment.ID).Updates(&comment)
	if result.Error != nil {
		fmt.Printf("Error while updating comment %v\n", result.Error)
	}
	return result.Error
}

func DeleteComment(commentID string) error {
	result := postgre.DB.Where("id = ?", commentID).Delete(&Comment{})
	if result.Error != nil {
		fmt.Printf("Error while deleting comment %v\n", result.Error)
	}
	return result.Error
}

func GetAllCommentsByCustomerID(customerID string) (*[]Comment, error) {
	comments := []Comment{}
	result := postgre.DB.Where("customer_id = ?", customerID).Find(&comments)
	if result.Error != nil {
		fmt.Printf("Error while searching products : %v", result.Error)
		return nil, result.Error
	}
	return &comments, result.Error
}
