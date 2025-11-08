package model

type Transaction struct {
	Timestamp   int64   `json:"timestamp"`
	Name        string  `json:"name"`
	Type        string  `json:"type"` // CREDIT or DEBIT
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"` // SUCCESS, FAILED, PENDING
	Description string  `json:"description"`
}
