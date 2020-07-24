package card

import (
	"sort"
	"time"
)

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []*Transaction
}

type Transaction struct {
	Id     string
	Amount int64
	Date   int64
	Mcc    string
	Status string
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SumByMCC(transactions []*Transaction, mcc []string) int64 {
	total := int64(0)
	for _, transaction := range transactions {
		for _, example := range mcc {
			if example == transaction.Mcc {
				total += transaction.Amount
			}
		}
	}
	return total
}

func SortTransaction(transaction []*Transaction) []*Transaction {
	sort.SliceStable(transaction, func(i, j int) bool {
		return transaction[i].Amount > transaction[j].Amount
	})
	return transaction
}

//task 2
type TransactionTest struct {
	Id     string
	Amount int64
	Date   time.Time
}

func Sum(transactions []*TransactionTest) int64 {
	total := int64(0)
	for _, transaction := range transactions {
		total += transaction.Amount
	}
	return total
}
