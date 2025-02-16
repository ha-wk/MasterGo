package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Balance float64
	mu      sync.Mutex
}

func(ba *BankAccount) Deposit(amount float64){
	ba.mu.Lock()
	defer ba.mu.Unlock()

	if amount<=0{
		fmt.Println("Invalid AMOUNT,please enter valid ammount")
		return
	}
    
	ba.Balance+=amount
	fmt.Println("Amount succesfully deposited to your account")
	fmt.Printf("Deposited: %.2f | New Balance: %.2f\n", amount, ba.Balance)

}

func(ba *BankAccount) Withdraw(amount float64){
	ba.mu.Lock()
	defer ba.mu.Unlock()

	if amount<=0{
		fmt.Println("Invalid AMOUNT,please enter valid ammount")
		return
	} else if amount>ba.Balance{
       fmt.Println("Insufficient AMOUNT,please enter lesser ammount")
		return
	}
    
	ba.Balance-=amount
	fmt.Println("Amount succesfully withdrawed from your account")
	fmt.Printf("Withdrawed: %.2f | New Balance: %.2f\n", amount, ba.Balance)

}

func(ba *BankAccount) ShowBalance() float64{
	
	ba.mu.Lock()
	defer ba.mu.Unlock()

	return ba.Balance
}

func main(){
	var UserBankAccount BankAccount

	UserBankAccount.Deposit(34)
	UserBankAccount.Deposit(45)
	UserBankAccount.Withdraw(70)
	UserBankAccount.ShowBalance()
	UserBankAccount.Withdraw(50)
	UserBankAccount.Withdraw(7)
}