package main

func main() {

}

type DivisorError struct {
	Dividend int
	Divisor  int
	Msg      string
}

func (e *DivisorError) Error() string {
	return e.Msg
}
