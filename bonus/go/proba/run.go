package proba

type Duck struct {
	Esp        float64
	StdDev     float64
	TB5M       int
	TB5S       int
	TB9M       int
	TB9S       int
	PercentOne float64
	PercentTwo float64
}

func Run(a float64) *Duck {
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

	return &res
}
