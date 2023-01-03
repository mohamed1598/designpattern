package factory

import "fmt"

type CashPM struct{}

func (pm *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n",amount)
}