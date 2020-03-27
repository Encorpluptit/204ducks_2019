package main

import (
	argv "./argv"
	proba "./proba"
	"fmt"
	"os"
)

func main() {
	av := os.Args[1:]

	a := argv.Check(av)
	fmt.Println(a)
	interval := proba.IntervalCreate(100)
	//disp.PrintInterval(interval)
	esp := proba.Esperance(a, interval, 100)
	fmt.Println(esp)
	stdDev := proba.StdDev(a, esp, 100, interval)
	fmt.Println(stdDev)
}
