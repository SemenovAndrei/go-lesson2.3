package card

import (
	"errors"
	"fmt"
	"sort"
	"sync"
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

func PrintMonthlyTransaction(transactions map[int]map[time.Month][]*TransactionTest) error {
	if len(transactions) == 0 {
		fmt.Println(errors.New("транзакций не было"))
		return errors.New("транзакций не было")
	}

	for i := range transactions {
		fmt.Println()
		fmt.Println("год   :", i)
		for j := range transactions[i] {
			fmt.Println("месяц :", j)
			SumConcurrently(transactions[i][j], 5)
		}
	}
	return nil
}

func SumConcurrently(transactions []*TransactionTest, goroutines int) int64 {
	if goroutines > len(transactions) {
		goroutines = len(transactions)
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	total := int64(0)
	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			mu.Lock()
			sum := sum(part)
			total += sum
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	partLast := len(transactions) % goroutines
	if partLast != 0 {
		total += sum(transactions[len(transactions)-partLast:])
	}
	fmt.Println("сумма :", total)
	return total
}

func sum(transactions []*TransactionTest) int64 {
	result := int64(0)
	for i := 0; i < len(transactions); i++ {
		result += transactions[i].Amount
	}
	return result
}
