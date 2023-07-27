package Employees

type Employee struct {
	Name          string `json:"name"`
	Id            int    `json:"id"`
	Address       string `json:"address"`
	Specification string `json:"specification"`
}
