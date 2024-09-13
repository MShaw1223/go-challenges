package objects

import "fmt"

type BankAccount struct {
	AccountNum int
	Balance float32
}

func (this * BankAccount) See_Balance() {
	fmt.Printf("Balance for ID - %d: %f\n",this.AccountNum, this.Balance)
}

func (this * BankAccount) Deposit(dep_amount float32){
	this.Balance += dep_amount
	fmt.Printf("Deposit of £%f to account ID - %d - successful\n",dep_amount,this.AccountNum)
}

func (this * BankAccount) Withdraw(wdrw_amount float32){
	this.Balance -= wdrw_amount
	fmt.Printf("Withdrawal of £%f to account ID - %d - successful\n", wdrw_amount, this.AccountNum)
}
