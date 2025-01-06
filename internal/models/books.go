package models

import "time"

type Book struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	Author    string    `json:"author"`
	Synopsis  string    `json:"synopsis"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (b *Book) TableName() string {
	return "books"
}

func (b *Book) GetEntity() *Book {
	return b
}
