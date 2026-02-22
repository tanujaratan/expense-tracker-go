# System Design – Expense Tracker API

## Objective

Build a REST API using Go that allows users to record shared expenses and automatically calculate settlements between participants.

## Architecture

Client
↓
REST API Server (Go net/http)
↓
In-memory storage (Slices / Maps)

This lightweight architecture was chosen because the project is a prototype and does not require persistent storage.

## Components

### API Layer

Handles HTTP requests and responses.

Endpoints:

* POST /users
* POST /expense
* GET /settle

### Data Models

User
Expense
Transaction

These are defined in `models.go`.

### Business Logic

The settlement logic is implemented in `settlement.go`.

Steps:

1. Calculate total expenses
2. Split equally among users
3. Compute each user’s balance
4. Identify debtors and creditors
5. Generate transactions to settle balances

### Algorithm Used

A greedy settlement algorithm is used to minimize the number of payments required to settle debts.

## Future Improvements

* Add database (PostgreSQL / MongoDB)
* Authentication
* Group support
* UI frontend
* Docker deployment
