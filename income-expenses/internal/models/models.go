package models

type CreateIncomeExpensesRequest struct {
    Category string  `json:"category"`
    Currency string  `json:"currency"`
    Amount   float32 `json:"amount"`
    Date     string  `json:"date"`
    Type     string  `json:"type"`
}

type CreateIncomeExpensesResponse struct {
    Message       string `json:"message"`
    TransactionID string `json:"transactionid"`
}

type InfoResponse struct {
    TransactionID string  `json:"transactionid"`
    Type          string  `json:"type"`
    Category      string  `json:"category"`
    Currency      string  `json:"currency"`
    Amount        float32 `json:"amount"`
    Date          string  `json:"date"`
}