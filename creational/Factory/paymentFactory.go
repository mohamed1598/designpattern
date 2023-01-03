package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPm),nil
	}
	return nil, fmt.Errorf("payment Method %d not recognized", m)
}
