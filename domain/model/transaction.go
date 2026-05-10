package model

import "fmt"

type Transaction struct {
	Timestamp   int64   `json:"timestamp"`
	Name        string  `json:"name"`
	Type        string  `json:"type"` // CREDIT or DEBIT
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"` // SUCCESS, FAILED, PENDING
	Description string  `json:"description"`
}

func (t Transaction) Validate() error {
	if t.Timestamp == 0 {
		return fmt.Errorf("timestamp is empty")
	}
	if t.Name == "" {
		return fmt.Errorf("name is empty")
	}
	if t.Type == "" {
		return fmt.Errorf("type is empty")
	}
	if t.Amount == 0 {
		return fmt.Errorf("amount is empty")
	}
	if t.Status == "" {
		return fmt.Errorf("status is empty")
	}
	if t.Description == "" {
		return fmt.Errorf("description is empty")
	}
	return nil
}