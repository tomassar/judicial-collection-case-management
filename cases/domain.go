package cases

import "time"

type Case struct {
	ID           int       // Identificador único de la causa
	DebtorName   string    // Nombre del deudor
	Amount       int       // Cantidad adeudada
	CreationDate time.Time // Fecha de creación de la causa
	Status       string    // Estado actual del caso (por ejemplo: en proceso, resuelto, archivado, etc.)
	Documents    []string  // Lista de documentos relacionados con la causa
	// Otros atributos que podrían ser útiles:
	// - AttorneyName: Nombre del abogado a cargo del caso
	// - Court: Tribunal al que pertenece el caso
	// - NextHearingDate: Fecha de la próxima audiencia
	// - Notes: Notas o comentarios adicionales sobre la causa
}
