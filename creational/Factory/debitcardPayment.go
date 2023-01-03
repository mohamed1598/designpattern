package factory

import "fmt"

type DebitCardPM struct{}

func (pm *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n",amount)
}