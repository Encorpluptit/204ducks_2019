package tests

import (
	"../argv"
	"log"
	"os"
	"testing"
)

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
				"Aruments [%v]) incorrect, err is [%v] (Expected [%v]", table.arg, err, table.err)
		}
		if res != table.exp {
			t.Errorf(
				"Aruments [%v]) incorrect, err is [%v] (Expected [%v]", table.arg, err, table.err)
		}
	}
}
