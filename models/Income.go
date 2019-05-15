package models

type Income struct {
	ID     string  `json:"id,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	Hint   string  `json:"hint,omitempty"`
	Tags   string  `json:"tags,omitempty"`
	Date   string  `json:"date,omitempty"`
}

func Store(obj Income) bool {
	//TODO save Incomes to persistance storage
	return true
}

func NewIncome(amount float32, hint string, tags string) Income {
	i := Income{ID: "1", Amount: amount, Hint: hint, Tags: tags}
	Store(i)
	return i
}
