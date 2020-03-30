package argv

import (
	display "../display"
	"strconv"
)

var ExitError float64 = 84

func Check(av []string) float64 {
	if len(av) != 1 {
		return display.Help("Wrong number of arguments", ExitError)
	}
	a, err := strconv.ParseFloat(av[0], 64)
	if err != nil {
		return display.Help("Wrong argument", ExitError)
	}
	if a < 0 || a > 2.5 {
		return display.Help("Argument must be between 0 and 2.5", ExitError)
	}
	return a
}
