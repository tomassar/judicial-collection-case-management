package cases

import "time"

type CreateCaseReq struct {
	DebtorName string    `json:"debtor_name"`
	Status     string    `json:"status"`
	Role       string    `json:"role"`
	Date       time.Time `json:"date"`
	Subject    string    `json:"subject"`
	Court      string    `json:"court"`
	LawyerID   uint      `json:"lawyer_id"`
}
