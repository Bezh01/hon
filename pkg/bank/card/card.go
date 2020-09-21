package card

import (
	"bank/pkg/bank/types"
)

// Total func
func Total(cards []types.Card) types.Money {
	var sum types.Money = types.Money(0)
	for _, card := range cards {
		sum += card.Balance
	}
	return sum
}