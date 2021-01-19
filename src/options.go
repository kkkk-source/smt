package smt

import "flag"

var (
    S sFlag
    E = flag.Int("e", 0, eUsage)
    T = flag.Int("t", 0, tUsage)
    F = flag.Int("f", 0, fUsage)
    C = flag.Int("c", 0, cUsage)
    W = flag.Int("w", 0, wUsage)
)

const (
    eUsage = `
-e  APP_ID
    Execute APP_ID demostration:
        1 Echo Server
        2 Clock Server
        3 Disk Usage
        4 Load Animation
`
    cUsage = `
-c
    It follows the -e flag. Use it when you want to execute in a sequential fashion the Echo Server and Clock Server demostrations.
        1 Sequential fashion
        2 Concurrent fashion
`
    tUsage = `
-t  TCP_CLIENT_ID
    Execute TCP_CLIENT_ID:
        1 Read-only TCP client
        2 Read-write TCP client
`
    sUsage = `
-s  SIMULATION_ID
    Execute SIMULATION_ID:
        1 Financial Lack Race Condition Simulation
        2 No Single Machine Word Race Condition Simulation
`
    fUsage = `
-f  CORRECT_SIMULATION_ID
    Execute one simulations of corrects concurrent functions:
        1 Avoid Race Condition Second Way 
        2 Avoid Race Condition Third  Way
`
    wUsage = `
-w
     It Follows the -s or -f flags and avoid displaying the simulation information. Use it when you want to measure the execution time of those simulations.
`
)

type sFlag []string

func (s *sFlag) String() string {
    return "string representation"
}

func (s *sFlag) Set(value string) error {
    *s = append(*s, value)
    return nil
}

func init() {
    flag.Var(&S, "s", sUsage)
}
