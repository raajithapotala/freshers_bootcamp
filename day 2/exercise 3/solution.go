package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()

	fmt.Println(amount, "amount is being deposited")
	balance += amount

	mutex.Unlock()
}

func Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()

	if amount > balance {
		fmt.Println("Insufficient Balance")
	} else {
		fmt.Println(amount, "amount is withdrawn")
		balance -= amount
	}

	mutex.Unlock()
}

func main() {
	balance = 500
	var wg sync.WaitGroup
	wg.Add(3)
	go Withdraw(300, &wg)
	go Deposit(100, &wg)
	go Withdraw(400, &wg)

	wg.Wait()

	fmt.Println("The current balance is ", balance)
}
