package argv

import (
	display "../display"
	"strconv"
)

var ExitError int

func Check(av []string) float64 {
	if len(av) != 1 {
		display.Help("Wrong number of arguments", ExitError)
	}
	a, err := strconv.ParseFloat(av[0], 64)
	if err != nil || 0 > a || a > 2.5 {
		display.Help("Wrong argument", ExitError)
	}
	return a
}
