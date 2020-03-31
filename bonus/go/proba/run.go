package proba

import (
	"fmt"
	"math"
)

type Duck struct {
	Esp        float64
	StdDev     float64
	TB5M       int64
	TB5S       int64
	TB9M       int64
	TB9S       int64
	PercentOne float64
	PercentTwo float64
}

func Run(a float64) {
	interval := IntervalCreate(100)
	intervalTime := IntervalTimeBack(100)

	esp := Esperance(a, interval, 100)
	res := Duck{
		Esp:        esp,
		StdDev:     StdDev(a, esp, 100, interval),
		PercentOne: PercentBack(a, 1),
		PercentTwo: PercentBack(a, 2),
	}
	res.TB5M, res.TB5S = DivMod(TimeBack(a, 50., intervalTime)*60, 60)
	res.TB9M, res.TB9S = DivMod(TimeBack(a, 99., intervalTime)*60, 60)

	Result(&res)
}

func Result(res *Duck) {
	fmt.Printf("Average return time: %0.0fm %0.02ds\n",
		math.Round(res.Esp*60)/60, int(math.RoundToEven(res.Esp*60))%60)
	fmt.Printf("Standard deviation: %.3f\n", res.StdDev)
	fmt.Printf("Time after which 50%% of the ducks are back: %dm %02ds\n", res.TB5M, res.TB5S)
	fmt.Printf("Time after which 99%% of the ducks are back: %dm %02ds\n", res.TB9M, res.TB9S)
	fmt.Printf("Percentage of ducks back after 1 minute: %.1f%%\n", res.PercentOne)
	fmt.Printf("Percentage of ducks back after 2 minutes: %.1f%%\n", res.PercentTwo)
}
