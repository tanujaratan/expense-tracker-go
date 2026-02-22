# Expense Tracker & Bill Settlement API (Golang)

## Overview

This project is a simple REST API built using Go that helps groups track shared expenses and automatically calculate how debts should be settled. The system works similarly to apps like Splitwise.

The goal is to reduce the number of transactions required when multiple people share expenses.

---

## Problem

When friends or roommates share expenses (trips, food, rent, etc.), calculating who owes whom becomes confusing and time-consuming.

---

## Solution

This API records shared expenses and computes the optimal settlement so that debts are cleared with the minimum number of payments.

---

## Tech Stack

* Go (Golang)
* Standard Library (`net/http`)
* JSON APIs
* In-memory data structures

---

## Project Structure

expense-tracker-go
│
├── main.go → REST API server
├── models.go → Data models
├── settlement.go → Debt settlement algorithm
├── design.md → System design explanation
├── prompts.md → AI prompts used
└── go.mod → Go module configuration

---

## API Endpoints

### Create User

POST /users

Example Request
{
"id": 1,
"name": "A"
}

---

### Add Expense

POST /expense

Example Request
{
"paid_by": 1,
"amount": 900
}

---

### Get Settlement

GET /settle

Example Response
[
{
"from": "B",
"to": "A",
"amount": 300
},
{
"from": "C",
"to": "A",
"amount": 300
}
]

---

## How the Algorithm Works

1. Calculate the total expense.
2. Divide it equally among all users.
3. Determine how much each user overpaid or owes.
4. Match debtors with creditors.
5. Generate the minimum set of transactions required to settle balances.

---

## Running the Project

1. Install Go
2. Clone the repository
3. Run:

go run .

Server runs on:
http://localhost:8080

---

## Example Scenario

If three users share an expense and one user pays ₹900:

Each person's share = ₹300.

Result:

* B pays A ₹300
* C pays A ₹300

---

## AI Assistance

AI tools were used to help with project structuring and implementation.
All prompts used are included in `prompts.md`.

---

## Author

Tanuja Ratan
