package main

import (
	argv "./argv"
	disp "./display"
	proba "./proba"
	"os"
)

func main() {
	av := os.Args[1:]
	a := argv.Check(av)
	if a == argv.ExitError {
		os.Exit(int(argv.ExitError))
	}
	interval := proba.IntervalCreate(100)
	esp := proba.Esperance(a, interval, 100)
	stdDev := proba.StdDev(a, esp, 100, interval)
	disp.Result(esp, stdDev, a)
}
