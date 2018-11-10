package entities

//Income struct represent Income object
type Income struct {
	ID     int64       `json:"id"`
	Amount float32     `json:"amount"`
	Hint   string      `json:"hint,omitempty"`
	Tags   []IncomeTag `json:"tags,omitempty"`
}
