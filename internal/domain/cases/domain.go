package cases

import (
	"time"

	"gorm.io/gorm"
)

type Case struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	DebtorName string `json:"debtor_name"` // Nombre del deudor
	Status     string `json:"status"`      // Estado actual del caso (por ejemplo: en proceso, resuelto, archivado, etc.)
	//Documents    []string  // Lista de documentos relacionados con la causa
	// Otros atributos que podrían ser útiles:
	// - AttorneyName: Nombre del abogado a cargo del caso
	// - Court: Tribunal al que pertenece el caso
	// - NextHearingDate: Fecha de la próxima audiencia
	// - Notes: Notas o comentarios adicionales sobre la causa
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
