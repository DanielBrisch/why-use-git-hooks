package main

import "git-hooks/transaction"

func main() {
	transactions := []transaction.Transaction{
		{ID: "T001", Amount: 120.50},
		{ID: "T002", Amount: 80.00},
		{ID: "T003", Amount: 300.75},
		{ID: "T004", Amount: 50.00},
	}

	results := transaction.ProcessTransactions(transactions)
	transaction.PrintReport(results)
}
