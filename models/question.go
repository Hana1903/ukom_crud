// package models

// import (
// 	"time"
// )

// type Question struct {
// 	ID            int64     `gorm:"primaryKey;autoIncrement"`
// 	IDPackage     int64     `gorm:"not null"`
// 	Question      string    `gorm:"type:text;not null"`
// 	Answer        string    `gorm:"type:varchar(255);default:null"`
// 	CorrectAnswer string    `gorm:"type:varchar(255);default:null"`
// 		CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }

package models

import (
	"time"
)

type Question struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	IDPackage     int64     `gorm:"not null"`
	Question      string    `gorm:"type:text;not null"`
	Answer        string    `gorm:"type:varchar(255);default:null"`
	CorrectAnswer string    `gorm:"type:varchar(255);default:null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"` // otomatis diisi saat pertama kali record dibuat
	UpdatedAt     time.Time `gorm:"autoUpdateTime"` // otomatis diupdate setiap kali record diubah
}
