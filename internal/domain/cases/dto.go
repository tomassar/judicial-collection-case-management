package cases

type CreateCaseReq struct {
	DebtorName string `json:"debtor_name"`
	Status     string `json:"status"`
	LawyerID   uint   `json:"lawyer_id"`
}
