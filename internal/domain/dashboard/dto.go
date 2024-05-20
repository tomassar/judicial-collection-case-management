package dashboard

import "github.com/tomassar/judicial-collection-case-management/internal/domain/cases"

type DisplayDataRes struct {
	Cases []*cases.Case
}
