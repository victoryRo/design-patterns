package factory

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func ExampleCashPM_Pay() {
	pay, _ := GetPaymentMethod(Cash)

	msg := pay.Pay(21.04)
	fmt.Println(msg)
	// Output: 21.04 paid using cash
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExisten(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Error("A payment method with ID 20 must return error")
	}
	t.Log("LOG:", err)
}
