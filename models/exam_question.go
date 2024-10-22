package models

import (
	"time"
)

type ExamQuestion struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	ExamID      int64     `gorm:"not null"`  // Foreign key referencing Exam
	QuestionID   int64     `gorm:"not null"`  // Foreign key referencing Question
	UserAnswer  string    `gorm:"type:text;not null"` // User's answer to the question
	CreatedAt   time.Time `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
