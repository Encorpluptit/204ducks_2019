package argv

import (
	"fmt"
	"strconv"
)

var ExitError = 84

var ExitSuccess = 0

func Check(av []string) (float64, *int) {
	if len(av) != 1 {
		return help("Wrong number of arguments", &ExitError)
	}
	if av[0] == "-h" {
		return help("", &ExitSuccess)
	}

	a, err := strconv.ParseFloat(av[0], 64)
	if err != nil {
		return help("Wrong argument", &ExitError)
	}

	if a < 0 || a > 2.5 {
		return help("Argument must be between 0 and 2.5", &ExitError)
	}

	return a, nil
}

func help(str string, ret *int) (float64, *int) {
	fmt.Println(str)
	return 0., ret
}
