package models

type IncomeTag struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
}

func Store(obj IncomeTag) bool {
	//TODO save IncomeTags to persistance storage
	return true;
}

func NewIncomeTag(name float32, description string, aliases []string) IncomeTag {
	i := IncomeTag{Id: "1", Name: name, Description: description, Aliases: aliases}
	Store(i)
	return i
}