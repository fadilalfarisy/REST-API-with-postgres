package models

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NameBook  string    `gorm:"not null;type:varchar(100)"  json:"name_book"`
	Author    string    `gorm:"not null;type:varchar(100)"  json:"author"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}
