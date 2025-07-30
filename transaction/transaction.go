package transaction

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Transaction struct {
	ID     string
	Amount float64
	Status string
}

func ProcessTransactions(transactions []Transaction) []Transaction {
	var wg sync.WaitGroup
	results := make([]Transaction, len(transactions))

	for i, t := range transactions {
		wg.Add(1)
		go func(i int, t Transaction) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)

			if rand.Intn(100) > 15 {
				t.Status = "Approved"
			} else {
				t.Status = "Refused"
			}
			results[i] = t
		}(i, t)
	}

	wg.Wait()
	return results
}

func PrintReport(transactions []Transaction) {
	for _, t := range transactions {
		fmt.Printf("ID: %s | Value: R$%.2f | Status: %s\n", t.ID, t.Amount, t.Status)
	}
}
