package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	temp := PaymentSource(
		[]types.Card{
			{
				PAN:     "5058 xxxx xxxx 8888",
				Balance: 25_000_00,  Active: true,
			}, 
			{
				PAN:     "5058 xxxx xxxx 8888",
				Balance: 30_000_00,
				Active:  false,
			},
			{
				PAN:     "5058 xxxx xxxx 8888",
				Balance: 40_000_00,
				Active:  true,
			},
		},
	)
	for _, buff := range temp {
		fmt.Println(buff.Number)
	}

	// Output: 
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 8888

}