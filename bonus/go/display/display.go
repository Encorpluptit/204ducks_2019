package display

import (
	"../proba"
	"fmt"
	"math"
)

func divmod(numerator, denominator float64) (quotient, remainder int64) {
	quotient = int64(numerator / denominator)
	remainder = int64(numerator) % int64(denominator)
	return
}

func Help(str string, ret float64) float64 {
	fmt.Println(str)
	return ret
}


func Result(esp float64, stdDev float64, a float64) {
	fmt.Printf("Average return time: %0.0fm %0.02ds\n", math.Round(esp*60)/60, int(math.RoundToEven(esp*60))%60)
	fmt.Printf("Standard deviation: %.3f\n", stdDev)
	interval := proba.IntervalTimeBack(100)
	q, r := divmod(proba.TimeBack(a,50., interval) * 60, 60)
	fmt.Printf("Time after which 50%% of the ducks are back: %dm %02ds\n", q, r)
	q, r = divmod(proba.TimeBack(a,99., interval) * 60, 60)
	fmt.Printf("Time after which 99%% of the ducks are back: %dm %02ds\n", q, r)
	fmt.Printf("Percentage of ducks back after 1 minute: %.1f%%\n", proba.PercentBack(a, 1))
	fmt.Printf("Percentage of ducks back after 2 minutes: %.1f%%\n", proba.PercentBack(a, 2))
}
