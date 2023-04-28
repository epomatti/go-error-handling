package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	errorTypeAssertion()

	errorBuiltinAs()

	errorSwitch()

	directComparison()
}

// Custom Error
type PaymentError struct {
	Reference string
	Amount    float64
	Message   string
	Timestamp time.Time
}

func newPaymentError(ref string, amt float64, msg string) *PaymentError {
	return &PaymentError{
		Reference: ref,
		Amount:    amt,
		Message:   msg,
		Timestamp: time.Now(),
	}
}

func (e *PaymentError) Error() string {
	ts := e.Timestamp.Format("2006-01-02 15:04:05")

	return fmt.Sprintf("Payment Error Ref: %s, Amt: %.2f, Msg: %s, Time: %s", e.Reference, e.Amount, e.Message, ts)
}

var ErrInvalidPaymentType = errors.New("INVALID PAYMENT TYPE") // Sentinel error

func processPayment(ref string, amt float64) error {
	if ref == "" {
		return ErrInvalidPaymentType
	}

	if amt > 100.0 {
		return newPaymentError(ref, amt, "Insufficient Funds")
	}

	return nil
}

// Error detection with assertion

func errorTypeAssertion() {
	err := processPayment("ABC", 120.00)
	if errAssert, ok := err.(*PaymentError); ok {
		fmt.Println("Is a paymento error")
		fmt.Println(errAssert)
	} else {
		fmt.Println("Is not a payment error")
		fmt.Println(errAssert)
	}
}

// Detection with errors.As function
func errorBuiltinAs() {
	err := processPayment("ABC", 120.00)
	var pmtErr *PaymentError
	if errors.As(err, &pmtErr) {
		fmt.Println("Is a paymento error")
		fmt.Println(pmtErr)
	} else {
		fmt.Println("Is not a payment error")
		fmt.Println(err)
	}
}

// Detection with switch
func errorSwitch() {
	if err := processPayment("ABC", 120.00); err != nil {
		switch e := err.(type) {
		case *PaymentError:
			fmt.Println("Is a payment error")
			fmt.Println(e)
		default:
			fmt.Println("Is not a payment error")
			fmt.Println(e)
		}
	}
}

// Direct comparison
// Does not perform a text comparison of the message, so it is important to return and compare the same sentinel error
func directComparison() {
	if err := processPayment("", 120.00); err != nil {
		switch err {
		case ErrInvalidPaymentType:
			fmt.Println("Is an invalid payment reference error")
			fmt.Println(err)
		default:
			fmt.Println("Is not an invalid payment reference error")
			fmt.Println(err)
		}
	}
}
