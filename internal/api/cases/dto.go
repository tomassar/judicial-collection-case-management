package cases

type CreateCaseReq struct {
	DebtorName string `json:"debtor_name"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
}
