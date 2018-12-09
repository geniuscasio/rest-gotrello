package entities

import "time"

//Income struct represent Income object
type Income struct {
	ID     int64       `json:"id"`
	Amount float32     `json:"amount"`
	Date   time.Time   `json:"date"`
	Hint   string      `json:"hint"`
	Tags   []IncomeTag `json:"tags,omitempty"`
}
