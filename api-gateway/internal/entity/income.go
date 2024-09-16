package entity

type (
	CreateIncomeExpensesRequest struct {
		Category string  `json:"category"`
		Currency string  `json:"currency"`
		Amount   float32 `json:"amount"`
		Date     string  `json:"date"`
		Type     string  `json:"type"`
		UserID   string  `json:"-"`
	}

	CreateIncomeExpensesResponse struct {
		Message       string `json:"message"`
		TransactionID string `json:"-"`
	}

	GetInfoRequest struct {
		UserID string `json:"-"`
	}

	InfoResponse struct {
		TransactionID string  `json:"transaction_id"`
		Type          string  `json:"type"`
		Category      string  `json:"category"`
		Currency      string  `json:"currency"`
		Amount        float32 `json:"amount"`
		Date          string  `json:"date"`
		UserID        string  `json:"user_id"`
	}

	GetInfoResponse struct {
		Info []InfoResponse `json:"info"`
	}
)
