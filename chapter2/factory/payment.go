package factory

import (
	"fmt"
)

// PaymentMethod defines a way of paying in the shop. This factory method returns
// objects that implements this interface
type PaymentMethod interface {
	Pay(amount float32) string
}

// --------------------------------------

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}

// --------------------------------------

type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card", amount)
}

// --------------------------------------

type NewDebitCardPM struct{}

func (c *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (new)", amount)
}

// --------------------------------------

// Our current implemented Payment methods are described here
const (
	Cash      = 1
	DebitCard = 2
)

// CreatePaymentMethod returns a pointer to a PaymentMethod object or an error
// if the method is not registered. We used "new" operator to return the pointer
// but we could also used &Type{} althought new makes it more readable for
// newcomers could be confusing
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(NewDebitCardPM), nil
	default:
		return nil, fmt.Errorf("Payment method %d not recognize", m)
	}
}
