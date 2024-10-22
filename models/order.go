// package models

// import (
// 	"time"
// )

// type Order struct {
// 	ID           int64     `gorm:"primaryKey;autoIncrement"`
// 	UserID       int64     `gorm:"not null"`                        
// 	PackageID    int64     `gorm:"not null"`                        
// 	PaymentStatus string   `gorm:"type:varchar(255);not null"`      
// 	TotalPrice   float64   `gorm:"type:decimal(10,2);not null"`      
// 	CreatedAt    time.Time `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP"`
// 	UpdatedAt    time.Time `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP"`
// }

package models

import "time"

type Order struct {
	ID             int64      `gorm:"primaryKey" json:"id"`
	UserID         int64      `json:"user_id"`
	PackageID      int64      `json:"package_id"`
	PaymentStatus   string    `gorm:"type:varchar(255)" json:"payment_status"`
	TotalPrice      float64   `json:"total_price"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
