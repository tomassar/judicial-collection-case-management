package cases

import "time"

type CreateCaseReq struct {
	DebtorName string     `json:"debtor_name"`
	Status     string     `json:"status"`
	Role       string     `json:"role"`
	Date       CustomTime `json:"date"`
	Subject    string     `json:"subject"`
	Court      string     `json:"court"`
	LawyerID   uint       `json:"lawyer_id"`
}

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]                   // Remove quotes around the time string
	t, err := time.Parse("2006-01-02", s) // Adjust layout according to your time format
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) ToTime() time.Time {
	return time.Time(ct)
}
