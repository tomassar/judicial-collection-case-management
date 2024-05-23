package cases

import (
	"time"

	"gorm.io/gorm"
)

type Case struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	LawyerID   uint           `json:"lawyer_id"`
	DebtorName string         `json:"debtor_name"` // Nombre del deudor
	Status     string         `json:"status"`      // Estado actual del caso (por ejemplo: en proceso, resuelto, archivado, etc.)
	Role       string         `json:"rol"`
	Date       time.Time      `json:"date"`
	Subject    string         `json:"subject"` // caratulado
	Court      string         `json:"court"`   // tribunal
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
