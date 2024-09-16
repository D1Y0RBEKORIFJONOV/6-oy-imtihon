package entity

type (
	IncomeExpenseResponse struct {
		TotalIncome   float32 `json:"total_income"`
		TotalExpenses float32 `json:"total_expenses"`
		NetSavings    float32 `json:"net_savings"`
	}

	SpendingRequest struct {
		UserID string `json:"user_id"`
	}

	SpendingResponse struct {
		Category   string  `json:"category"`
		TotalSpent float32 `json:"total_spent"`
	}

	ListSpendingResponse struct {
		Spent []SpendingResponse `json:"spent"`
	}

	FromTillRequest struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		UserID    string `json:"user_id"`
	}

	FromTillResponse struct {
		TotalAmount float32 `json:"total_amount"`
	}
)
