package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/i-hit/go-lesson2.3.git/pkg/card"
)

func main() {
	master := &card.Card{
		Id:       1,
		Issuer:   "MasterCard",
		Balance:  100_000_00,
		Currency: "RUB",
		Number:   "0001",
		Transactions: []*card.Transaction{
			{
				Id:     "1",
				Amount: 101,
				Date:   time.Now().Unix(),
				Mcc:    "5011",
				Status: "В обработке",
			},
			{
				Id:     "2",
				Amount: 10,
				Date:   time.Now().Unix(),
				Mcc:    "5012",
				Status: "В обработке",
			},
		},
	}
	master.Transactions = append(master.Transactions,
		&card.Transaction{Id: "3", Amount: 10000, Date: time.Now().Unix(), Mcc: "5011", Status: "В обработке"})
	card.AddTransaction(master,
		&card.Transaction{Id: "4", Amount: 1000, Date: time.Now().Unix(), Mcc: "5011", Status: "Выполнено"})
	card.AddTransaction(master,
		&card.Transaction{Id: "5", Amount: 102, Date: time.Now().Unix(), Mcc: "5011", Status: "Выполнено"})

	var mcc []string
	mcc = append(mcc, "5011", "5014", "5015")
	fmt.Println(mcc)
	fmt.Println(card.TranslateMCC(master.Transactions[0].Mcc))
	fmt.Println()

	// sort
	fmt.Println("sort")
	card.SortTransaction(master.Transactions)

	for _, i := range master.Transactions {
		fmt.Println(i)
	}

	fmt.Println()

	//	task 2
	fmt.Println("task 2")
	transaction := make([]*card.TransactionTest, 365)
	var cnt int64 = 1
	date := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := range transaction {
		transaction[i] = &card.TransactionTest{Id: strconv.FormatInt(cnt, 10), Amount: 2,
			Date: date.AddDate(0, 0, int(cnt))}
		cnt++
	}

	transactionByMonth := make(map[time.Month][]*card.TransactionTest)

	for i := range transaction {
		transactionByMonth[transaction[i].Date.Month()] = append(transactionByMonth[transaction[i].Date.Month()], transaction[i])
	}

	transactionSliceByMonth := make([][]*card.TransactionTest, 12, 12)

	// test - 0 transaction in month
	// delete transaction in month
	transactionByMonth[time.Month(3)] = nil

	for i := time.Month(1); i <= time.Month(12); i++ {
		month := i
		transactionSliceByMonth[month - 1] = transactionByMonth[month]
	}

	card.SumConcurrently(transactionSliceByMonth, 12)
	fmt.Println()

}
