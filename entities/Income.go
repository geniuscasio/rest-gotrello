package entities

//Income struct represent Income object
type Income struct {
	ID     int64   `json:"id"`
	Amount float32 `json:"amount"`
	Date   string  `json:"date"`
	Hint   string  `json:"hint"`
	Tags   string  `json:"tags,omitempty"`
}
