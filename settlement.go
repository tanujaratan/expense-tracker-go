package main

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func CalculateSettlement(users []User, expenses []Expense) []Transaction {

	balance := make(map[int]float64)

	total := 0.0

	for _, e := range expenses {
		total += e.Amount
		balance[e.PaidBy] += e.Amount
	}

	split := total / float64(len(users))

	for _, u := range users {
		balance[u.ID] -= split
	}

	var debtors []int
	var creditors []int

	for id, bal := range balance {
		if bal < 0 {
			debtors = append(debtors, id)
		} else if bal > 0 {
			creditors = append(creditors, id)
		}
	}

	var transactions []Transaction

	for len(debtors) > 0 && len(creditors) > 0 {

		d := debtors[0]
		c := creditors[0]

		amount := min(-balance[d], balance[c])

		transactions = append(transactions, Transaction{
			From:   getUserName(d, users),
			To:     getUserName(c, users),
			Amount: amount,
		})

		balance[d] += amount
		balance[c] -= amount

		if balance[d] == 0 {
			debtors = debtors[1:]
		}

		if balance[c] == 0 {
			creditors = creditors[1:]
		}
	}

	return transactions
}

func getUserName(id int, users []User) string {
	for _, u := range users {
		if u.ID == id {
			return u.Name
		}
	}
	return ""
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}