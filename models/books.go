package models

import (
	"gorm.io/gorm"
	"time"
)

type Jobs struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	JobTitle  *string   `json"jobtitle"`
	Company   *string   `json"title"`
	ApplyLink *string   `json"publisher"`
	JobDesc   *string   `json"jobdesc"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	endDate   *string   `json"endDate"`
	salary    *string   `json"salary"`
	active    uint      `json"active"`
}

// ss
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Jobs{})
	return err
}
