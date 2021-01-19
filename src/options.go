package smt

import "flag"

var (
	SFlag sFlag
	EFlag = flag.Int("e", 0, eUsage)
	TFlag = flag.Int("t", 0, tUsage)
	FFlag = flag.Int("f", 0, fUsage)
	CFlag = flag.Int("c", 0, cUsage)
	WFlag = flag.Int("w", 0, wUsage)
)

const (
	eUsage = `
Execute APP_ID demostration:

    1 Echo Server
    2 Clock Server
    3 Disk Usage
    4 Load Animation
`
	cUsage = `
It follows the -e flag. Use it when you want to execute in a sequential 
or concurrent fashion the Echo Server and Clock Server demostrations.

    1 Sequential fashion
    2 Concurrent fashion
`
	tUsage = `
Execute TCP_CLIENT_ID:

    2 Read-write TCP client
    1 Read-only TCP client
`
	sUsage = `
Execute SIMULATION_ID:

    1 Financial Lack Race Condition Simulation
    2 No Single Machine Word Race Condition Simulation
`
	fUsage = `
Execute one simulations of corrects concurrent functions:

    1 Avoid Race Condition Second Way 
    2 Avoid Race Condition Third  Way
`
	wUsage = `
 It Follows the -s or -f flags and avoid displaying the simulation information.
 Use it when you want to measure the execution time of those simulations.
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
	flag.Var(&SFlag, "s", sUsage)
}
