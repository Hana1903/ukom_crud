package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateOfBirth time.Time

const dateLayout = "2006-01-02"

// UnmarshalJSON for custom date type
func (d *DateOfBirth) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Remove the surrounding quotes
	t, err := time.Parse(dateLayout, str)
	if err != nil {
		return fmt.Errorf("could not parse date: %v", err)
	}
	*d = DateOfBirth(t)
	return nil
}

// MarshalJSON for custom date type
func (d DateOfBirth) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(d).Format(dateLayout))), nil
}

// ToTime converts DateOfBirth to time.Time
func (d DateOfBirth) ToTime() time.Time {
	return time.Time(d)
}

// Value method for saving DateOfBirth to database
func (d DateOfBirth) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Scan method for reading DateOfBirth from database
func (d *DateOfBirth) Scan(value interface{}) error {
	if value == nil {
		*d = DateOfBirth(time.Time{})
		return nil
	}

	switch t := value.(type) {
	case time.Time:
		*d = DateOfBirth(t)
		return nil
	default:
		return fmt.Errorf("cannot convert %v to DateOfBirth", value)
	}
}

type User struct {
	ID                    uint        `gorm:"primaryKey"`
	Name                  string      `gorm:"type:varchar(255)"`
	Email                 string      `gorm:"uniqueIndex;type:varchar(255)"`
	Password              string      `gorm:"type:varchar(255)"`
	DateOfBirth           DateOfBirth `gorm:"type:date"`  
	Gender                string      `gorm:"type:varchar(20)"`
	PhoneNumber           string      `gorm:"type:varchar(20)"`
	EducationalInstitution string     `gorm:"type:varchar(255)"`
	Profession            string      `gorm:"type:varchar(255)"`
	Address               string      `gorm:"type:text"`
	Province              string      `gorm:"type:varchar(255)"`
	City                  string      `gorm:"type:varchar(255)"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
