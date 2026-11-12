
package models

import "time"

type Applicant struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(150);uniqueIndex;not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(20)"`
	Position  string    `json:"position" gorm:"type:varchar(100);not null"`
	Stage     string    `json:"stage" gorm:"type:varchar(20);not null;default:applied"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Applicant) TableName() string {
	return "applicants"
}