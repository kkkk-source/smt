package smt

import (
	"fmt"
	"os"
	"sync"
)

const financialLackRaceConditionSimulationInfo = `
 RACE CONDITION ISSUE SIMULATION
 _______________________________

+-{ Definition }--------------------------------------------------------------------------------------------+
|                                                                                                           |
| A race condition is a situation in which the program does not give the correct result for some            |
| interleavings of the operations of multiple threads. Race conditions are pernicious because they may      |
| remain latent in a program and appear infrequently, perhaps only under heavy load or when using certain   |
| compilers, platforms, or architectures. This makes them hard to reproduce and diagnose.                   |
|                                                                                                           |
+-{ Context }-----------------------------------------------------------------------------------------------+
|                                                                                                           |
| It is traditional to explain the seriousness of race conditions through the metaphor of financial         |
| loss, so we’ll consider a simple bank account program.                                                    |
|                                                                                                           |
| Let's suppose the following situation: Alice deposits %d, and Bob %d. There is a particular outcome,      
| in which Bob's deposit occurs in the middle of Alice's deposit, after the balance has been read but       |
| before it has been updated, causing Bob's transaction to disappear. This is because Alice's deposit       |
| operation is really a sequence of two operations, a read and a write. What is that special outcome?       |
|                                                                                                           |
+-{ Outcome }-----------------------------------------------------------------------------------------------+
|                                                                                                           |
| The specia outcome is %d. The expected outcome is %d but we got %d so we say, the Bob's deposit (%d)      
| was lost in heaven. Don't get lost in heaven. The number of attemps that were taken to get the            |
| special outcome were %d.                                                                                  
|                                                                                                           |
+-{ Function }----------------------------------------------------------------------------------------------+
|                                                                                                           |
| This was the executed function which gave us the previous special outcome. The function name is           |
| FinancialLackRaceConditionSimulation and always return the special outcome because of race condition. It  |
| also returns the number of attemps that were taken to get that special outcome. It takes two argument a,  |
| and b. Where a, and b are the amounts that are going to be deposited into the same bank account which     |
| always starts at 0.                                                                                       |
|                                                                                                           |
| func financialLackRaceConditionSimulation(a, b int) (int, int) {                                          |
|   var want, attemps int                                                                                   |
|   restoreBalance()                                                                                        |
|                                                                                                           |
|   want = a + b                                                                                            |
|   attemps = 0                                                                                             |
|   for true {                                                                                              |
|     var wg sync.WaitGroup                                                                                 |
|                                                                                                           |
|     wg.Add(2)                                                                                             |
|     go func() {                                                                                           |
|       setDeposit(a) <-- Race Condition here                                                               |
|       wg.Done()                                                                                           |
|     }()                                                                                                   |
|     go func() {                                                                                           |
|       setDeposit(b) <-- Race Condition here                                                               |
|       wg.Done()                                                                                           |
|     }()                                                                                                   |
|     wg.Wait()                                                                                             |
|     attemps++                                                                                             |
|     if got := getBalance(); got != want {                                                                 |
|       return got, attemps                                                                                 |
|     }                                                                                                     |
|     restoreBalance()                                                                                      |
|   }                                                                                                       |
|   return 0, 0                                                                                             |
| }                                                                                                         |
|                                                                                                           |
+-{ Critical Section }--------------------------------------------------------------------------------------+
|                                                                                                           |
| This was the section responsible for the special outcome.                                                 |
|                                                                                                           |
| func setDeposit(amount int) {                                                                             |
|   balance = balance + amount  <-- Critical Section                                                        |
| }                                                                                                         |
|                                                                                                           |
+-----------------------------------------------------------------------------------------------------------+
`

const noSingleMachineWordRaceConditionSimulationInfo = `
 RACE CONDITION ISSUE SIMULATION
 _______________________________

+-{ Definition }--------------------------------------------------------------------------------------------+
|                                                                                                           |
| A race condition is a situation in which the program does not give the correct result for some            |
| interleavings of the operations of multiple threads. Race conditions are pernicious because they may      |
| remain latent in a program and appear infrequently, perhaps only under heavy load or when using certain   |
| compilers, platforms, or architectures. This makes them hard to reproduce and diagnose.                   |
|                                                                                                           |
+-{ Context }-----------------------------------------------------------------------------------------------+
|                                                                                                           |
| Things get even messier if the data race involves a variable of a type that is larger than a single       |
| machine word, such an interface, a string, or a slice: the pointer, the length, and the capacity.         |
|                                                                                                           |
|                                                                                                           |
| --A slice, is a dynamically-sized, flexible view into the elements of an array--                          |
|                                                                                                           |
+-{ Function }----------------------------------------------------------------------------------------------+
|                                                                                                           |
| This function updates concurretly two slices of different lengths. The value of x in the final statement  |
| is not defined; it could be nil, or a slice of length 10, or a slice of length 1,000,000. But recall that |
| there are three parts to a slice: the pointer, the length, and the capacity. If the pointer comes from    |
| the first call to make and the length comes from the second, x would be a chimera, a slice whose nominal  |
| length is 1,000,000 but whose underlying array has only 10 elements. In this eventuality, storing to      |
| element 999,999 would clobber an arbitrary faraway memory location, with consequences that are imposible  |
| to predict and hard to debug and localize. This semantic minefield is called undefined behavior and is    |
| well known to C programmers.                                                                              |
|                                                                                                           |
| func noSingleMachineWordRaceConditionSimulation() {                                                       |
|   var x []int                                                                                             |
|   go func() {                                                                                             |
|     x = make([]int, 10)                                                                                   |
|   }()                                                                                                     |
|   go func() {                                                                                             |
|     x = make([]int, 1000000)                                                                              |
|   }()                                                                                                     |
|   x[999999] = 1 -> NOTE: undefined behavior; memory corruption possible!                                  |
| }                                                                                                         |
|                                                                                                           |
+-----------------------------------------------------------------------------------------------------------+
`

const avoidRaceCondition = `
                                          --Before going through this section make sure you have executed 
                                                    the Financial Lack Race Condition Simulation already.--

 AVOIDING RACE CONDITION
 _______________________

+-{ Context }-----------------------------------------------------------------------------------------------+
|                                                                                                           |
| Did you remember the FinancialLackRaceConditionSimulation? Well let's us put it and the critical section  |
| just here (here we go again):                                                                             |
|                                                                                                           |
| func financialLackRaceConditionSimulation(a, b int) (int, int) {                                          |
|   var want, attemps int                                                                                   |
|   restoreBalance()                                                                                        |
|                                                                                                           |
|   want = a + b                                                                                            |
|   attemps = 0                                                                                             |
|   for true {                                                                                              |
|     var wg sync.WaitGroup                                                                                 |
|                                                                                                           |
|     wg.Add(2)                                                                                             |
|     go func() {                                                                                           |
|       setDeposit(a) <-- Race Condition here                                                               |
|       wg.Done()                                                                                           |
|     }()                                                                                                   |
|     go func() {                                                                                           |
|       setDeposit(b) <-- Race Condition here                                                               |
|       wg.Done()                                                                                           |
|     }()                                                                                                   |
|     wg.Wait()                                                                                             |
|     attemps++                                                                                             |
|     if got := getBalance(); got != want {                                                                 |
|       return got, attemps                                                                                 |
|     }                                                                                                     |
|     restoreBalance()                                                                                      |
|   }                                                                                                       |
|   return 0, 0                                                                                             |
| }                                                                                                         |
|                                                                                                           |
| func setDeposit(amount int) {                                                                             |
|   balance = balance + amount  <-- Critical Section                                                        |
| }                                                                                                         |
|                                                                                                           |
+-{ Fix it }------------------------------------------------------------------------------------------------+
|                                                                                                           |
| There are three ways to avoid data race.                                                                  |
|                                                                                                           |
| 1) The first way is not to write the variable. If instead we initialize with all necessary entries before |
| creating threads and never modify it again, then any number of threads may safely call the related        |
| function. But what can we do if we need to modify the entries? Let's read a little bit more.              |
|                                                                                                           |
| 2) The second way to avoid data race is to avoid accesing the variable from multiple threads and confined |
| it to a single thread. Since other threads cannot access the variable directly, they must use a channel   |
| to send the confinnig thread a request to query or update the variable. This is what is meant by the Go   |
| mantra "Do not communicate by sharing memory; instad, share memory by communicating.".                    |
|                                                                                                           |
| Now let's fix our previous FinancialLackRaceConditionSimulation code with this approach.                  |
|                                                                                                           |
| func avoidDataRaceSecondWay(a, b int) int {                                                               |
|   go Teller()                                                                                             |
|   var wg sync.WaitGroup                                                                                   |
|                                                                                                           |
|   wg.Add(2)                                                                                               |
|   go func() {                                                                                             |
|     setDeposits(a)                                                                                        |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|   go func() {                                                                                             |
|     setDeposits(b)                                                                                        |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|                                                                                                           |
|   wg.Wait()                                                                                               |
|   got := getBalances()                                                                                    |
|   restoreBalance()                                                                                        |
|   return got                                                                                              |
| }                                                                                                         |
|                                                                                                           |
| func setDeposits(amount int) {                                                                            |
|   deposits <- amount                                                                                      |
| }                                                                                                         |
|                                                                                                           |
| func getBalances() int {                                                                                  |
|   return <-balances                                                                                       |
| }                                                                                                         |
|                                                                                                           |
| func teller() {                                                                                           |
|   var balance int                                                                                         |
|   for {                                                                                                   |
|     select {                                                                                              |
|     case amount := <-deposits:                                                                            |
|       balance += amount                                                                                   |
|     case balances <- balance:                                                                             |
|     }                                                                                                     |
|   }                                                                                                       |
| }                                                                                                         |
|                                                                                                           |
| 3) The third way to avoid a data race is to allow many threads to access the variable, but only one at a  |
| time. This approach is known as mutual exclusion and is the subject of the next section.                  |
|                                                                                                           |
| func avoidDataRaceThirdWay(a, b int) int {                                                                |
|   var wg sync.WaitGroup                                                                                   |
|   var mu sync.Mutex                                                                                       |
|                                                                                                           |
|   restoreBalance()                                                                                        |
|   wg.Add(2)                                                                                               |
|   go func() {                                                                                             |
|     mu.Lock()     <-- lock                                                                                |
|     setDeposit(a)  <-- critical section                                                                   |
|     mu.Unlock()   <-- unlock                                                                              |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|   go func() {                                                                                             |
|     mu.Lock()     <-- lock                                                                                |
|     setDeposit(b)    <-- critical section                                                                 |
|     mu.Unlock()   <-- unlock                                                                              |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|   wg.Wait()                                                                                               |
|   got := getBalance()                                                                                     |
|   restoreBalance()                                                                                        |
|   return got                                                                                              |
| }                                                                                                         |
|                                                                                                           |
+-----------------------------------------------------------------------------------------------------------+
`

const avoidRaceConditionSimulation = `
                                          --Before going through this section make sure you have executed 
                                                    the Financial Lack Race Condition Simulation already.--
 AVOID RACE CONDITION SIMULATION
 _______________________________

+-{ Definition }--------------------------------------------------------------------------------------------+
|                                                                                                           |
| Recall that a race condition is a situation in which the program does not give the correct result for     |
| some interleavings of the operations of multiple threads. Race conditions are pernicious because they may |
| remain latent in a program and appear infrequently, perhaps only under heavy load or when using certain   |
| compilers, platforms, or architectures. This makes them hard to reproduce and diagnose.                   |
|                                                                                                           |
+-{ Context }-----------------------------------------------------------------------------------------------+
|                                                                                                           |
| In the previous Financial Lack Race Condition Simulation we have an issue. If Alice deposits %d, and 
| Bob %d. when Alice or Bob wants to read their bank account, they could get an outcome like this %d 
| instead of %d, which is the correct outcome.
|                                                                                                           |
+-{ Outcomes }----------------------------------------------------------------------------------------------+
|                                                                                                           |
| Let's run the simulation again but this time we are going to use the adquired knowledge of Avoid Race     |
| Condition section.                                                                                        |
|                                                                                                           |
+-----------------------------------------------------------------------------------------------------------+
|
| Executing the following function, which referes to the second way of avoiding race condition from [ Avoid 
| Race Condition ] section, you would get %d, when Alice deposits %d and Bob %d. So this approach works 
| correctly.                                                                                                |
|                                                                                                           |
| func avoidDataRaceSecondWay(a, b int) int {                                                               |
|   go Teller()                                                                                             |
|   var wg sync.WaitGroup                                                                                   |
|                                                                                                           |
|   wg.Add(2)                                                                                               |
|   go func() {                                                                                             |
|     setDeposits(a)                                                                                        |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|   go func() {                                                                                             |
|     setDeposits(b)                                                                                        |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|                                                                                                           |
|   wg.Wait()                                                                                               |
|   got := getBalances()                                                                                    |
|   restoreBalance()                                                                                        |
|   return got                                                                                              |
| }                                                                                                         |
|                                                                                                           |
| func setDeposits(amount int) {                                                                            |
|   deposits <- amount                                                                                      |
| }                                                                                                         |
|                                                                                                           |
| func getBalances() int {                                                                                  |
|   return <-balances                                                                                       |
| }                                                                                                         |
|                                                                                                           |
| func Teller() {                                                                                           |
|   var balance int                                                                                         |
|   for {                                                                                                   |
|     select {                                                                                              |
|     case amount := <-deposits:                                                                            |
|       balance += amount                                                                                   |
|     case balances <- balance:                                                                             |
|     }                                                                                                     |
|   }                                                                                                       |
| }                                                                                                         |
|                                                                                                           |
| Executing the following function, which referes to the third way of avoiding race condition from [ Avoid  | 
| Race Condition ] section, you would get %d, when Alice deposits %d and Bob %d. So this approach works     
| correctly as well.                                                                                        |
|                                                                                                           |
| func avoidDataRaceThirdWay(a, b int) int {                                                                |
|   var wg sync.WaitGroup                                                                                   |
|   var mu sync.Mutex                                                                                       |
|                                                                                                           |
|   restoreBalance()                                                                                        |
|   wg.Add(2)                                                                                               |
|   go func() {                                                                                             |
|     mu.Lock()                                                                                             |
|     setDeposit(a)                                                                                         |
|     mu.Unlock()                                                                                           |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|   go func() {                                                                                             |
|     mu.Lock()                                                                                             |
|     setDeposit(b)                                                                                         |
|     mu.Unlock()                                                                                           |
|     wg.Done()                                                                                             |
|   }()                                                                                                     |
|   wg.Wait()                                                                                               |
|   got := getBalance()                                                                                     |
|   restoreBalance()                                                                                        |
|   return got                                                                                              |
| }                                                                                                         |
|                                                                                                           |
+-----------------------------------------------------------------------------------------------------------+
`

func FinancialLackSimulation(alice, bob int) {
	want := alice + bob
	got, attemps := financialLackRaceConditionSimulation(alice, bob)
	fmt.Fprintf(os.Stderr, financialLackRaceConditionSimulationInfo, alice, bob, got, want, got, bob, attemps)
}

func NoSingleMachineWordSimulation() {
	fmt.Fprintf(os.Stderr, noSingleMachineWordRaceConditionSimulationInfo)
}

func AvoidDataRace(alice, bob int) {
	fmt.Fprintf(os.Stderr, avoidRaceCondition)
	want := alice + bob
	gotA := avoidDataRaceSecondWay(alice, bob)
	gotB := avoidDataRaceThirdWay(alice, bob)
	gotC, _ := financialLackRaceConditionSimulation(alice, bob)
	fmt.Fprintf(os.Stderr, avoidRaceConditionSimulation, alice, bob, gotC, want, gotA, alice, bob, gotB, alice, bob)
}

var (
	balance  int
	deposits = make(chan int) // send amount to deposits
	balances = make(chan int) // receive balance
)

func restoreBalance() {
	balance = 0
}

func setDeposit(amount int) {
	// critical section
	balance = balance + amount
}

func getBalance() int {
	return balance
}

func setDeposits(amount int) {
	deposits <- amount
}

func getBalances() int {
	return <-balances
}

// Monitor goroutine
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

// FinancialLackRaceConditionSimulation always return the special outcome because
// of race condition and the number of attemps that were taken to get that special
// outcome. It takes two argument a, and b. Where a and b are the amounts that are
// going to be deposited into the same bank account which always starts at 0.
func financialLackRaceConditionSimulation(a, b int) (int, int) {
	var want, attemps int
	restoreBalance()

	want = a + b
	attemps = 0
	for true {
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			setDeposit(a)
			wg.Done()
		}()
		go func() {
			setDeposit(b)
			wg.Done()
		}()
		wg.Wait()
		attemps++
		if got := getBalance(); got != want {
			return got, attemps
		}
		restoreBalance()
	}
	return 0, 0
}

func avoidDataRaceSecondWay(a, b int) int {
	go teller() // start the monitor
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		setDeposits(a)
		wg.Done()
	}()
	go func() {
		setDeposits(b)
		wg.Done()
	}()

	wg.Wait()
	got := getBalances()
	restoreBalance()
	return got
}

func avoidDataRaceThirdWay(a, b int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	restoreBalance()
	wg.Add(2)
	go func() {
		mu.Lock()
		setDeposit(a)
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		mu.Lock()
		setDeposit(b)
		mu.Unlock()
		wg.Done()
	}()
	wg.Wait()
	got := getBalance()
	restoreBalance()
	return got
}
