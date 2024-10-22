package models

import "time"

type Package struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	NamePackage  string     `gorm:"type:varchar(255)" json:"name_package"`
	Description  string     `gorm:"type:text" json:"description"`
	Price        float64    `gorm:"type:decimal" json:"price"`
	DurationExam int        `gorm:"type:integer" json:"duration_exam"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Question    []Question `gorm:"foreignKey:IDPackage" json:"questions"`
	// Payment      string     `gorm:"type:varchar(255)" json:"payment"`
}
