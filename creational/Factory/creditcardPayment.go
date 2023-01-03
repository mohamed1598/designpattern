package factory

import "fmt"

type CreditCardPm struct{}

func (pm CreditCardPm) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit card\n",amount)
}
