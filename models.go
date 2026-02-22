package main

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Expense struct {
	PaidBy int     `json:"paid_by"`
	Amount float64 `json:"amount"`
}