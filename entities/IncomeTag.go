package entities

//IncomeTag struct represent IncIncome object
type IncomeTag struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description"`
	Aliases     []string `json:"aliases"`
}
