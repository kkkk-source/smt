package smt

import (
	"flag"
	"os"
	"strconv"
)

func Run() {
	var port string
	var cway bool
	flag.Parse()

	switch *CFlag {
	case 1:
		cway = false
	default:
		cway = true
	}

	port = "8000"
	switch *TFlag {
	case 1:
		MakeProof(EchoServerProof, port)
	case 2:
		MakeProof(ClockServerProof, port)
	}

	switch *EFlag {
	case 1:
		MakeServer(EchoServer, port, cway)
	case 2:
		MakeServer(ClockServer, port, cway)
	case 3:
		DiskUsage([]string{"."})
		os.Exit(0)
	case 4:
		SpinnerAnimation()
		os.Exit(0)
	}

	SFlagLen := len(SFlag)
	if SFlagLen != 0 {
		switch SFlag[0] {
		case "2":
			NoSingleMachineWordSimulation()
		case "1":
			if SFlagLen == 3 {
				alice, err := strconv.Atoi(SFlag[1])
				if err != nil {
					os.Exit(1)
				}
				bob, err := strconv.Atoi(SFlag[2])
				if err != nil {
					os.Exit(1)
				}
				FinancialLackSimulation(alice, bob)
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	}

	FFlagLen := len(FFlag)
	if FFlagLen != 0 {
		switch FFlag[0] {
		case "1":
			if FFlagLen == 3 {
				alice, err := strconv.Atoi(FFlag[1])
				if err != nil {
					os.Exit(1)
				}
				bob, err := strconv.Atoi(FFlag[2])
				if err != nil {
					os.Exit(1)
				}
				AvoidDataRace(alice, bob)
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	}
	os.Exit(0)
}
