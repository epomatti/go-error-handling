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

func Test_Calc(t *testing.T) {
	tests := []struct {
		X   int
		Y   int
		Res int
		Err string
	}{
		{
			X: 1, Y: -1, Res: 0, Err: "y > 0",
		},
		{
			X: 1, Y: 1, Res: 0, Err: "x > y",
		},
		{
			X: 5, Y: 2, Res: 7, Err: "",
		},
	}

	for _, test := range tests {
		res, err := Calc(test.X, test.Y)

		if res != test.Res {
			t.Errorf("Result = %d, want %d", res, test.Res)
		}

		eMsg := ""
		if err != nil {
			eMsg = err.Error()
		}

		if eMsg != test.Err {
			t.Errorf("Err = '%s', want '%s'", eMsg, test.Err)
		}
	}
}
