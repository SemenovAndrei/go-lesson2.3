package card

import (
	"reflect"
	"testing"
	"time"
)



func TestSortTransaction(t *testing.T) {
	type args struct {
		transaction []*Transaction
	}
	
	
	
	tests := []struct {
		name string
		args args
		want []*Transaction
	}{
		{
			name: "first test",
			args: args{
				transaction: []*Transaction{
					{
						Id:     "1",
						Amount: 1,
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
					{
						Id:     "3",
						Amount: 101,
						Date:   time.Now().Unix(),
						Mcc:    "5012",
						Status: "В обработке",
					},
					{
						Id:     "3",
						Amount: 102,
						Date:   time.Now().Unix(),
						Mcc:    "5012",
						Status: "В обработке",
					},
					{
						Id:     "4",
						Amount: 3,
						Date:   time.Now().Unix(),
						Mcc:    "5012",
						Status: "В обработке",
					},
					{
						Id:     "5",
						Amount: 10000,
						Date:   time.Now().Unix(),
						Mcc:    "5012",
						Status: "В обработке",
					},
				},
			},
			want: []*Transaction{
				{
					Id:     "5",
					Amount: 10000,
					Date:   time.Now().Unix(),
					Mcc:    "5012",
					Status: "В обработке",
				},
				{
					Id:     "3",
					Amount: 102,
					Date:   time.Now().Unix(),
					Mcc:    "5012",
					Status: "В обработке",
				},
				{
					Id:     "3",
					Amount: 101,
					Date:   time.Now().Unix(),
					Mcc:    "5012",
					Status: "В обработке",
				},
				{
					Id:     "2",
					Amount: 10,
					Date:   time.Now().Unix(),
					Mcc:    "5012",
					Status: "В обработке",
				},
				{
					Id:     "4",
					Amount: 3,
					Date:   time.Now().Unix(),
					Mcc:    "5012",
					Status: "В обработке",
				},
				{
					Id:     "1",
					Amount: 1,
					Date:   time.Now().Unix(),
					Mcc:    "5011",
					Status: "В обработке",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortTransaction(tt.args.transaction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
