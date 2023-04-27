package main

import "testing"

// Erro test
func Test_NegativeY(t *testing.T) {
	res, err := Calc(1, -1)

	if res != 0 {
		t.Errorf("Result = %d, want 0", res)
	}

	if err == nil || err.Error() != "y > 0" {
		t.Errorf("Error = '%v', want 'y > 0'", err)
	}
}

// Error test
func Test_XLessThanY(t *testing.T) {
	res, err := Calc(1, 1)

	if res != 0 {
		t.Errorf("Result = %d, want 0", res)
	}

	if err == nil || err.Error() != "x > y" {
		t.Errorf("Error = '%v', want 'x > y'", err)
	}
}

// Result test
func Test_XBiggerThanY(t *testing.T) {
	res, err := Calc(5, 2)

	if res != 7 {
		t.Errorf("Result = %d, want 7", res)
	}

	if err != nil {
		t.Errorf("Error = '%v', want nil", err)
	}
}
