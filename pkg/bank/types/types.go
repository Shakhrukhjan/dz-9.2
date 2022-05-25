package types

const (
	TJS currency = "TJS"
	RUB currency = "RUB"
	USD currency = "USD"
)

type currency string

type Money int
  

type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
type Card struct {
	PAN     string
	Balance Money
	Active  bool
}
