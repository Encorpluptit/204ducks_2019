package main

import (
	argv "./argv"
	disp "./display"
	proba "./proba"
	"fmt"
	"os"
)

func main() {
	av := os.Args[1:]

	a := argv.Check(av)
	fmt.Println(a)
	interval := proba.IntervalCreate(100)
	esp := proba.Esperance(a, interval, 100)
	fmt.Println(esp)
	stdDev := proba.StdDev(a, esp, 100, interval)
	disp.Result(esp, stdDev, a)
	fmt.Println(stdDev)
}
