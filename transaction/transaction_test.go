package transaction

import "testing"

func TestProcessTransactions(t *testing.T) {
	transactions := []Transaction{
		{ID: "T001", Amount: 100},
		{ID: "T002", Amount: 50},
	}

	results := ProcessTransactions(transactions)

	if len(results) != len(transactions) {
		t.Errorf("Expected %d transactions, but received %d", len(transactions), len(results))
	}

	for _, r := range results {
		if r.Status != "Approved" && r.Status != "Refused" {
			t.Errorf("Invalid status for transaction %s: %s", r.ID, r.Status)
		}
	}
}
