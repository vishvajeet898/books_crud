package models

import (
	"gorm.io/gorm"
	"time"
)

type Jobs struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	JobTitle  *string   `json:"jobtitle"`
	Company   *string   `json:"title"`
	ApplyLink *string   `json:"publisher"`
	JobDesc   *string   `json:"jobdesc"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	EndDate   time.Time `json:"endDate"`
	Salary    *string   `json:"salary"`
	Active    bool      `json:"active"`
}

// ss
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Jobs{})
	return err
}
