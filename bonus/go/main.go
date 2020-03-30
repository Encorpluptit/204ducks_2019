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

	run(a)
}

func run(a float64) {
	interval := proba.IntervalCreate(100)
	intervalTime := proba.IntervalTimeBack(100)

	esp := proba.Esperance(a, interval, 100)
	res := disp.Duck{
		Esp:        esp,
		StdDev:     proba.StdDev(a, esp, 100, interval),
		PercentOne: proba.PercentBack(a, 1),
		PercentTwo: proba.PercentBack(a, 2),
	}
	res.TB5M, res.TB5S = proba.DivMod(proba.TimeBack(a, 50., intervalTime)*60, 60)
	res.TB9M, res.TB9S = proba.DivMod(proba.TimeBack(a, 99., intervalTime)*60, 60)

	disp.Result(&res)
}
