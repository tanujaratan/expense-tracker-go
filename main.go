package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users []User
var expenses []Expense

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)

	json.NewEncoder(w).Encode(user)
}

func addExpense(w http.ResponseWriter, r *http.Request) {
	var exp Expense
	json.NewDecoder(r.Body).Decode(&exp)

	expenses = append(expenses, exp)

	json.NewEncoder(w).Encode(exp)
}

func getSettlement(w http.ResponseWriter, r *http.Request) {

	result := CalculateSettlement(users, expenses)

	json.NewEncoder(w).Encode(result)
}

func main() {

	http.HandleFunc("/users", createUser)
	http.HandleFunc("/expense", addExpense)
	http.HandleFunc("/settle", getSettlement)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}