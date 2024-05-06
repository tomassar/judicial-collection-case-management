package cases

import (
	"time"

	"gorm.io/gorm"
)

type entityCommonAttrs struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type caseEntity struct {
	entityCommonAttrs
	DebtorName string `json:"debtor_name"` // Nombre del deudor
	Status     string `json:"status"`      // Estado actual del caso (por ejemplo: en proceso, resuelto, archivado, etc.)
	//Documents    []string  // Lista de documentos relacionados con la causa
	// Otros atributos que podrían ser útiles:
	// - AttorneyName: Nombre del abogado a cargo del caso
	// - Court: Tribunal al que pertenece el caso
	// - NextHearingDate: Fecha de la próxima audiencia
	// - Notes: Notas o comentarios adicionales sobre la causa
}

func (e *caseEntity) toModel() *caseModel {
	return &caseModel{
		Model:      entityCommonAttrsToGormModel(e.entityCommonAttrs),
		DebtorName: e.DebtorName,
		Status:     e.Status,
	}
}

func entityCommonAttrsToGormModel(e entityCommonAttrs) gorm.Model {
	return gorm.Model{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}
