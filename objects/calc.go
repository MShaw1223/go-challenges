package objects

// alter this so its just calc.add(6,9) etc so change any funcs not wanted to be exported to lowercase start

import (
	"errors"
)

type errs struct {
	InvalidOp error
	DivByZero error
}

var e = errs{
	DivByZero: errors.New("divide by 0 error"),
	InvalidOp: errors.New("invalid operation"),
}

type Calculator struct{}

func (Calculator) Add(a, b float64) (float64, error) {
	return a + b, nil
}

func (Calculator) Subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func (Calculator) Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, e.DivByZero
	}
	return a / b, nil
}

func (Calculator) Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func (Calculator) Describe(operation string) string {
	switch operation {
	case "add":
		return "Add: Join something to something else so as to increase the size, number or amount\n"
	case "subtract":
		return "Subtract: Take away a number or amount from another to calculate the difference\n"
	case "divide":
		return "Divide:Separate or be separated into parts\n"
	case "multiply":
		return "Multiply: Obtain from a number another which contains the first number a specified number of times\n"

	default:
		return e.InvalidOp.Error()
	}
}
