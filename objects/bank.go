package objects

import (
	"fmt"
	"math"
)

type BankAccount struct {
	AccountNum int
	Balance    float32
	InterestRt float32
}

func (this *BankAccount) See_Balance() {
	fmt.Printf("Balance for ID - %d: %f\n", this.AccountNum, this.Balance)
}

func (this *BankAccount) Deposit(dep_amount float32) {
	this.Balance += dep_amount
	fmt.Printf("Deposit of £%f to account ID - %d - successful\n", dep_amount, this.AccountNum)
}

func (this *BankAccount) Withdraw(wdrw_amount float32) {
	this.Balance -= wdrw_amount
	fmt.Printf("Withdrawal of £%f to account ID - %d - successful\n", wdrw_amount, this.AccountNum)
}
func (this *BankAccount) Returns(savings_dep_amount, years float32) {
	decimal_int_rt := this.InterestRt / 100
	returns := savings_dep_amount * float32(math.Pow(float64(1+decimal_int_rt), float64(years)))
	fmt.Printf("A deposit of £%f, will grow to £%f over 3 years", savings_dep_amount, returns)
}
