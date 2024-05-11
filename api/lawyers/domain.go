package lawyers

import (
	"time"

	"gorm.io/gorm"
)

type Lawyer struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	UserID        int            `json:"user_id"`
	FullName      string         `json:"full_name"`
	PhoneNumber   string         `json:"phone_number"`
	LicenseNumber string         `json:"license_number"`
	FirmName      string         `json:"firm_name"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
