package cases

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	DebtorName string `json:"debtor_name"` // Nombre del deudor
	Status     string `json:"status"`      // Estado actual del caso (por ejemplo: en proceso, resuelto, archivado, etc.)
	//Documents    []string  // Lista de documentos relacionados con la causa
	// Otros atributos que podrían ser útiles:
	// - AttorneyName: Nombre del abogado a cargo del caso
	// - Court: Tribunal al que pertenece el caso
	// - NextHearingDate: Fecha de la próxima audiencia
	// - Notes: Notas o comentarios adicionales sobre la causa
}

func (m *Case) toEntity() *caseEntity {
	return &caseEntity{
		entityCommonAttrs: gormModelToEntity(m.Model),
		DebtorName:        m.DebtorName,
	}
}

func caseModelsToEntities(models []*Case) (ent []*caseEntity) {
	for _, m := range models {
		ent = append(ent, m.toEntity())
	}

	return ent
}

func gormModelToEntity(m gorm.Model) entityCommonAttrs {
	return entityCommonAttrs{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}
