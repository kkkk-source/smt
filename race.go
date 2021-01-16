package main

import "sync"

var RaceConditionDescription = `A race condition is a situation in which the program does not give the correct result for 
some interleavings of the operations of multiple threads. Race conditions are pernicious because 
they may remain latent in a program and appear infrequently, perhaps only under heavy load or when 
using certain compilers, platforms, or architectures. This makes them hard to reproduce and diagnose.`

var NoSingleMachineWordStr = `Things get even messier if the data race involves a variable of a type that is larger than a single 
machine word, such an interface, a string, or a slice: the pointer, the length, and the capacity.
This code update concurretly two slices of different lengths:

func NoSingleMachineWord() {
    var x []int
    go func() { x = make([]int, 10) }()
    go func() { x = make([]int, 1000000) }()
    x[999999] = 1 -> NOTE: undefined behavior; memory corruption possible!
}

The value of x in the final statement is not defined; it could be nil, or a slice of length 10, or 
a slice of length 1,000,000. But recall that there are three parts to a slice: the pointer, the length, 
and the capacity. If the pointer comes from the first call to make and the length comes from the second, 
x would be a chimera, a slice whose nominal length is 1,000,000 but whose underlying array has only 10 
elements.  In this eventuality, storing to element 999,999 would clobber an arbitrary faraway memory 
location, with consequences that are imposible to predict and hard to debug and localize. This semantic
minefield is called undefined behavior and is well known to C programmers.
`

var DoFinancialLackStr = `
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
`

var CriticalSectionStr = `
func Deposit(amount int) {
	balance = balance + amount
}
`

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

func RestoreBalance() {
	balance = 0
}

var balance int

// Critical Section
func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

var (
	deposits = make(chan int) // send amount to deposits
	balances = make(chan int) // receive balance
)

func Deposits(amount int) {
	deposits <- amount
}

func Balances() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

var AvoidDataRaceSecondWayDescription = `The second way to avoid data race is to avoid accesing the variable from multiple threads and confined 
it to a single thread.  Since other threads cannot access the variable directly, they must use a channel 
to send the confinnig thread a request to query or update the variable. This is what is meant by the Go 
mantra "Do not communicate by sharing memory; instad, share memory by communicating."`

func AvoidDataRaceSecondWay(a, b int) int {
	go teller() // start the monitor
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		Deposits(a)
		wg.Done()
	}()
	go func() {
		Deposits(b)
		wg.Done()
	}()

	wg.Wait()
	got := Balances()
	RestoreBalance()
	return got
}

var AvoidDataRaceThirdWayDescription = ` The third way to avoid a data race is to allow many threads to access the variable, but only 
one at a time. This approach is known as mutual exclusion and is the subject of the next section.`

func AvoidDataRaceThirdWay(a, b int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	RestoreBalance()
	wg.Add(2)
	go func() {
		mu.Lock()
		Deposit(a)
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		mu.Lock()
		Deposit(b)
		mu.Unlock()
		wg.Done()
	}()
	wg.Wait()
	got := Balance()
	RestoreBalance()
	return got
}
