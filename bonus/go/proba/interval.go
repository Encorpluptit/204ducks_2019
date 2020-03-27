package proba

func IntervalCreate(end float64) []float64 {
	var interval []float64

	for i := 0.; i <= end; i += 0.001 {
		interval = append(interval, i)
	}
	return interval
}
