package main

import (
	"errors"
	"fmt"
	"time"
)

// Main
func main() {
	ErrorBuiltinAs()
}

// Error object types: sentinel, customer, wrapping

type error interface {
	Error() string
}

func DoSomething() error {
	return errors.New("something didn't work")
}

// The Go Way
func TheGoWay() {
	result, err := Divide(9, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("THE DIVISOR CANNOT BE ZERO")
	}
	return a / b, nil
}

// Complex error type
type DivisorError struct {
	Dividend int
	Divisor  int
	Msg      string
}

func (e *DivisorError) Error() string {
	return e.Msg
}

// Format errors
func FormattedError() error {
	return fmt.Errorf("user does not exist: %s", "john")
}

// Parsing dates
var format = "2006-01-02"
var value = "9999-15-31"

func ParseDate() {
	if _, err := time.Parse(format, value); err != nil {
		fmt.Println(err)
	}
}

// Panic
func Print(x, y int) {
	if y <= 0 {
		panic(fmt.Sprintf("%v", y))
	}

	fmt.Println("Result is", x/y)
}

// Recover
func Recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	Print(9, 0)
	fmt.Println("Completed") // This will not be called
}

// Testing
func Calc(x int, y int) (total int, err error) {
	if y < 0 {
		return 0, errors.New("y > 0")
	}

	if x <= y {
		return 0, errors.New("x > y")
	}

	return x + y, nil
}

// Custom Error
type PaymentError struct {
	Reference string
	Amount    float64
	Message   string
	Timestamp time.Time
}

func NewPaymentError(ref string, amt float64, msg string) *PaymentError {
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

func ProcessPayment(ref string, amt float64) error {
	if ref == "" {
		return ErrInvalidPaymentType
	}

	if amt > 100.0 {
		return NewPaymentError(ref, amt, "Insufficient Funds")
	}

	return nil
}

// Error detection with assertion

func ErrorTypeAssertion() {
	err := ProcessPayment("ABC", 120.00)
	if errAssert, ok := err.(*PaymentError); ok {
		fmt.Println("Is a paymento error")
		fmt.Println(errAssert)
	} else {
		fmt.Println("Is not a payment error")
		fmt.Println(errAssert)
	}
}

// Detection with errors.As function
func ErrorBuiltinAs() {
	err := ProcessPayment("ABC", 120.00)
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
func ErrorSwitch() {
	if err := ProcessPayment("ABC", 120.00); err != nil {
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
func DirectComparison() {
	if err := ProcessPayment("", 120.00); err != nil {
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
