package card

import (
	"bank/pkg/bank/types"
)

func PaymentSource(cards []types.Card) []types.PaymentSource {
	newPay := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance > 0 {

			newPay =  append(newPay, types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance})
		}
	}
	return newPay
}

