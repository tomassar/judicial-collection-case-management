package lawyers

import (
	"time"

	"gorm.io/gorm"
)

type Lawyer struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      int            `json:"user_id" gorm:"unique"`
	FullName    string         `json:"full_name"`
	PhoneNumber string         `json:"phone_number"`
	Rut         string         `json:"rut"`
	NationalID  string         `json:"national_id"` //rut
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
