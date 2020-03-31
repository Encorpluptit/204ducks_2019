package tests

import (
	"../argv"
	"../proba"
	"log"
	"math"
	"os"
	"testing"
)

type testDuck struct {
	EspM       int
	EspS       int
	stdDev     float64
	TB5M       int
	TB5S       int
	TB9M       int
	TB9S       int
	PercentOne float64
	PercentTwo float64
}

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

func TestArgv(t *testing.T) {
	defer quiet()()
	tables := []struct {
		arg []string
		exp float64
		err *int
	}{
		{[]string{"-0.5"}, 0., &argv.ExitError},
		{[]string{"-84"}, 0., &argv.ExitError},
		{[]string{"50"}, 0., &argv.ExitError},
		{[]string{"3"}, 0., &argv.ExitError},
		{[]string{"a"}, 0., &argv.ExitError},
		{[]string{"2a"}, 0., &argv.ExitError},
		{[]string{"3."}, 0., &argv.ExitError},
		{[]string{""}, 0., &argv.ExitError},
		{[]string{"3", "4"}, 0., &argv.ExitError},
		{[]string{"2."}, 2., nil},
		{[]string{"0.2"}, 0.2, nil},
		{[]string{"1.7"}, 1.7, nil},
		{[]string{"1.8"}, 1.8, nil},
		{[]string{"2.4"}, 2.4, nil},
		{[]string{"2.5"}, 2.5, nil},
		{[]string{"0"}, 0, nil},
	}

	for _, table := range tables {
		res, err := argv.Check(table.arg)
		if err != table.err {
			t.Errorf(
				"For argument(s) [%v]), err is [%v] (Expected [%v]\n", table.arg, err, table.err)
		}
		if res != table.exp {
			t.Errorf(
				"For argument(s) [%v]), res is [%v] (Expected [%v]\n", table.arg, res, table.exp)
		}
	}
}

func TestRunDuck(t *testing.T) {
	tables := []struct {
		a    float64
		duck testDuck
	}{
		{0.5, testDuck{
			EspM: 0, EspS: 56, stdDev: 0.808, TB5M: 0, TB5S: 43, TB9M: 3, TB9S: 57, PercentOne: 66.0, PercentTwo: 91.0},
		},
		{2.1, testDuck{
			EspM: 1, EspS: 32, stdDev: 1.126, TB5M: 1, TB5S: 16, TB9M: 5, TB9S: 21, PercentOne: 38.2, PercentTwo: 73.7},
		},
		{1.9, testDuck{
			EspM: 1, EspS: 28, stdDev: 1.109, TB5M: 1, TB5S: 11, TB9M: 5, TB9S: 15, PercentOne: 41.7, PercentTwo: 75.8},
		},
		{1.5, testDuck{
			EspM: 1, EspS: 19, stdDev: 1.059, TB5M: 1, TB5S: 1, TB9M: 5, TB9S: 0, PercentOne: 48.6, PercentTwo: 80.2},
		},
		{1.6, testDuck{
			EspM: 1, EspS: 21, stdDev: 1.074, TB5M: 1, TB5S: 4, TB9M: 5, TB9S: 4, PercentOne: 46.9, PercentTwo: 79.1},
		},
		{1.2, testDuck{
			EspM: 1, EspS: 12, stdDev: 1.005, TB5M: 0, TB5S: 55, TB9M: 4, TB9S: 47, PercentOne: 53.9, PercentTwo: 83.4},
		},
		{0.9, testDuck{
			EspM: 1, EspS: 5, stdDev: 0.935, TB5M: 0, TB5S: 49, TB9M: 4, TB9S: 30, PercentOne: 59.1, PercentTwo: 86.6},
		},
		{0.2, testDuck{
			EspM: 0, EspS: 50, stdDev: 0.676, TB5M: 0, TB5S: 39, TB9M: 3, TB9S: 16, PercentOne: 71.3, PercentTwo: 94.2},
		},
		{2.4, testDuck{
			EspM: 1, EspS: 39, stdDev: 1.141, TB5M: 1, TB5S: 23, TB9M: 5, TB9S: 28, PercentOne: 33.0, PercentTwo: 70.4},
		},
	}

	for _, table := range tables {
		RunTesting(t, table.a, &table.duck, proba.Run(table.a))
	}
}

func RunTesting(t *testing.T, a float64, exp *testDuck, res *proba.Duck) {
	espM, espS := int(math.Round(res.Esp*60)/60), int(math.RoundToEven(res.Esp*60))%60
	if espM != exp.EspM || espS != exp.EspS {
		t.Errorf(
			"For a = [%v], Esperance failed. Got [%v], [%v], (Expected [%v], [%v])\n",
			a, espM, espS, exp.EspM, exp.EspS)
	}
	if res.TB5M != exp.TB5M || res.TB5S != exp.TB5S {
		t.Errorf(
			"For a = [%v], Time Back (50%%) failed. Got [%v], [%v], (Expected [%v], [%v])\n",
			a, res.TB5M, res.TB5S, exp.TB5M, exp.TB5S)
	}
	if res.TB9M != exp.TB9M || res.TB9S != exp.TB9S {
		t.Errorf(
			"For a = [%v], Time Back (99%%) failed. Got [%v], [%v], (Expected [%v], [%v])\n",
			a, res.TB9M, res.TB9S, exp.TB9M, exp.TB9S)
	}
	if math.Round(res.PercentOne*10)/10 != exp.PercentOne {
		t.Errorf(
			"For a = [%v], Percent Back (1 min) failed. Got [%v] (Expected [%v])\n",
			a, res.PercentOne, exp.PercentOne)
	}
	if math.Round(res.PercentTwo*10)/10 != exp.PercentTwo {
		t.Errorf(
			"For a = [%v], Percent Back (2 mins) failed. Got [%v] (Expected [%v])\n",
			a, math.Round(res.PercentTwo*10)/10, exp.PercentTwo)
	}
}
