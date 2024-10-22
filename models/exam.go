package models

import "time"

type Exam struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int       `json:"order_id"`
	PackageID int       `json:"package_id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Score     float64   `json:"score"`
}