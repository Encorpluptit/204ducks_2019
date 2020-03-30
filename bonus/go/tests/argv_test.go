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
	}{
		{[]string{"-0.5"}, argv.ExitError},
		{[]string{"-84"}, argv.ExitError},
		{[]string{"50"}, argv.ExitError},
		{[]string{"3"}, argv.ExitError},
		{[]string{"a"}, argv.ExitError},
		{[]string{"2a"}, argv.ExitError},
		{[]string{"3."}, argv.ExitError},
		{[]string{""}, argv.ExitError},
		{[]string{"3", "4"}, argv.ExitError},
		{[]string{"2."}, 2.},
		{[]string{"0.2"}, 0.2},
		{[]string{"1.7"}, 1.7},
		{[]string{"1.8"}, 1.8},
		{[]string{"2.4"}, 2.4},
		{[]string{"2.5"}, 2.5},
		{[]string{"0"}, 0},
	}

	for _, table := range tables {
		res := argv.Check(table.arg)
		if res != table.exp {
			t.Errorf("Aruments [%v]) was incorrect", table.arg)
		}
	}
}
