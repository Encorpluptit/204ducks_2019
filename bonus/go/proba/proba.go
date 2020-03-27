package proba

import "math"

func ProbabilityDensity(a float64, t float64) float64 {
	return a*math.Exp(-t) + (4-3*a)*math.Exp(-2*t) + (2*a-4)*math.Exp(-4*t)
}

func Variance(esp float64, a float64, t float64) float64 {
	return math.Pow(t-esp, 2) * (ProbabilityDensity(a, t) / 10)
}

func Esperance(a float64, interval []float64, end float64) float64 {
	var result float64 = 0
	var i float64 = 0
	for idx := 0; i < end; i += 0.001 {
		result += interval[idx] * ProbabilityDensity(a, i) / 10
		idx += 1
	}
	return result / interval[len(interval)-1]
}

func StdDev(a float64, esp float64, end float64, interval []float64) float64 {
	var result float64 = 0
	var i float64 = 0
	for idx := 0; i < end; i += 0.001 {
		result += Variance(esp, a, i)
		idx += 1
	}
	return math.Sqrt(result / interval[len(interval)-1])
}
