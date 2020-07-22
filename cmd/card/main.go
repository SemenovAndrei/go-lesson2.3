package main

import (
	"fmt"
	"github.com/i-hit/go-lesson2.3.git/pkg/card"
	"time"
)

func main() {
	master := &card.Card{
		Id : 1,
		Issuer : "MasterCard",
		Balance : 100_000_00,
		Currency : "RUB",
		Number : "0001",
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
	fmt.Println(card.SumByMCC(master.Transactions, mcc))

	category := card.TranslateMCC(master.Transactions[0].Mcc)
	fmt.Println(category)
	fmt.Println()


	// sort
	card.SortTransaction(master.Transactions)

	for _, i := range master.Transactions {
		fmt.Println(i)
	}
}


