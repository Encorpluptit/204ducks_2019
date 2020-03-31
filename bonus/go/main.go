package main

import (
	argv "./argv"
	proba "./proba"
	"os"
)

func main() {
	av := os.Args[1:]
	a, err := argv.Check(av)

	if err != nil {
		os.Exit(*err)
	}

	proba.Run(a)
}
