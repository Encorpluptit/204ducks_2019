package display

import (
	proba "../proba"
	"fmt"
	"math"
	"os"
)

func Help(str string, ret int) {
	fmt.Println(str)
	os.Exit(ret)
}

func PrintInterval(s []float64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Result(esp float64, stdDev float64, a float64) {
	fmt.Printf("Average return time: %0.0fm %0.02ds\n", math.Round(esp*60)/60, int(math.RoundToEven(esp*60))%60)
	fmt.Printf("Standard deviation: %.3f\n", stdDev)
	//fmt.Printf("Time after which 50%% of the ducks are back: %dm %02ds\n",
	//	)
}
