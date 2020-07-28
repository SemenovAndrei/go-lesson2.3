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
	transaction := make([]*card.TransactionTest, 1100)
	cnt := 1
	date := time.Now()

	for i := range transaction {
		transaction[i] = &card.TransactionTest{Id: strconv.FormatInt(int64(cnt), 10), Amount: 10,
			Date: date.AddDate(0, 0, -cnt)}
		cnt++
	}

	transactionMap := make(map[int]map[time.Month][]*card.TransactionTest)

	for i := range transaction {
		year := transaction[i].Date.Year()
		month := transaction[i].Date.Month()

		if transactionMap[year] == nil {
			transactionMap[year] = map[time.Month][]*card.TransactionTest{}
		}

		transactionMap[year][month] = append(transactionMap[year][month], transaction[i])
	}

	_ = card.PrintMonthlyTransaction(transactionMap)
}
