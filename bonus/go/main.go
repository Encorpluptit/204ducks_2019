package main

import (
	argv "./argv"
	proba "./proba"
	"fmt"
	"math"
	"os"
)

func main() {
	av := os.Args[1:]
	a, err := argv.Check(av)

	if err != nil {
		os.Exit(*err)
	}

	printResult(proba.Run(a))
}

func printResult(res *proba.Duck) {
	fmt.Printf("Average return time: %0.0dm %0.02ds\n",
		int(math.Round(res.Esp*60)/60), int(math.RoundToEven(res.Esp*60))%60)
	fmt.Printf("Standard deviation: %.3f\n", res.StdDev)
	fmt.Printf("Time after which 50%% of the ducks are back: %dm %02ds\n", res.TB5M, res.TB5S)
	fmt.Printf("Time after which 99%% of the ducks are back: %dm %02ds\n", res.TB9M, res.TB9S)
	fmt.Printf("Percentage of ducks back after 1 minute: %.1f%%\n", res.PercentOne)
	fmt.Printf("Percentage of ducks back after 2 minutes: %.1f%%\n", res.PercentTwo)
}
