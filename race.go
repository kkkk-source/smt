package main

import "sync"

var RaceConditionDescription = `A race condition is a situation in which the program does not give the correct result for 
some interleavings of the operations of multiple threads. Race conditions are pernicious because 
they may remain latent in a program and appear infrequently, perhaps only under heavy load or when 
using certain compilers, platforms, or architectures. This makes them hard to reproduce and diagnose.`

// FinancialLack always return the special outcome because of race condition
// and the number of attemps that were taken to get that special outcome. It
// takes two argument a, and b. Where a and b are the amounts that are going
// to be deposited into the same bank account which always starts at 0.
func DoFinancialLack(a, b int) (int, int) {
	var want, attemps int
	RestoreBalance()

	want = a + b
	attemps = 0
	for true {
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			Deposit(a)
			wg.Done()
		}()
		go func() {
			Deposit(b)
			wg.Done()
		}()
		wg.Wait()
		attemps++
		if got := Balance(); got != want {
			return got, attemps
		}
		RestoreBalance()
	}
	return 0, 0
}

var balance int

func RestoreBalance() {
	balance = 0
}

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}
