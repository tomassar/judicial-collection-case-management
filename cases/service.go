package cases

import "time"

func GetCases() []Case {
	cases := []Case{
		{ID: 1, DebtorName: "John Doe", Amount: 1000.50, CreationDate: time.Now(), Status: "In progress", Documents: []string{"Contract", "Invoices"}},
		{ID: 2, DebtorName: "Jane Smith", Amount: 2000.75, CreationDate: time.Now(), Status: "Pending", Documents: []string{"Agreement", "Receipts"}},
	}

	return cases
}
