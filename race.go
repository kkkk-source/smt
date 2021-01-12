package main

import (
    "fmt"
    "sync"
)

var raceConditionDescription = `A race condition is a situation in which the program does not give the correct result for some interleavings of the operations of multiple threads. Race conditions are pernicious because they may remain latent in a program and appear infrequently, perhaps only under heavy load or when using certain compilers, platforms, orarchitectures. This makes them hard to reproduce and diagnose. 

It  is traditional to explain the seriousness of race conditions through the metaphor of financial loss, so we’ll consider a simple bank account program.`

var balance int

func Deposit(amount int) {
    balance = balance + amount
}

func Balance() int {
    return balance
}

func RaceConditionToPlayWith() {
    const a = 200
    const b = 100
    const want = a + b

    for (true) {
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
        if got := Balance(); got != want {
            fmt.Printf("got = %d, want = %d\n", got, want)
            return
        }
        balance = 0
    }
}
