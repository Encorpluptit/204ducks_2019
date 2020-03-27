package proba

import (
	"math"
)

func ProbabilityDensity(a float64, t float64) float64 {
	return a*math.Exp(-t) + (4-3*a)*math.Exp(-2*t) + (2*a-4)*math.Exp(-4*t)
}

func variance(esp float64, a float64, t float64) float64 {
	return math.Pow(t-esp, 2) * (ProbabilityDensity(a, t) / 10)
}

func Esperance(a float64, interval []float64, end float64) float64 {
	var result float64 = 0
	var i float64 = 0

	for idx := 0; i <= end; i += 0.001 {
		result += interval[idx] * ProbabilityDensity(a, i) / 10
		idx += 1
	}
	return result / interval[len(interval)-1]
}

func StdDev(a float64, esp float64, end float64, interval []float64) float64 {
	var result float64 = 0
	var i float64 = 0

	for idx := 0; i <= end; i += 0.001 {
		result += variance(esp, a, i)
		idx += 1
	}
	return math.Sqrt(result / interval[len(interval)-1])
}

func PercentBack(a float64, percent float64, end float64) float64 {
	var result float64 = 0
	var i float64 = 0
	end = end * percent
	interval := IntervalCreate(end)

	for idx := 0; i <= end; i += 0.001 {
		result += ProbabilityDensity(a, interval[idx]/end)
		idx += 1
	}
	return result / 10
}

//@staticmethod
//def percent_back(const: float, time_snd: int):
//return sum(Duck.probability_density(const, i / Duck.max) for i in range(time_snd * Duck.max)) / 10
//
//@staticmethod
//def time_back(a: float, p: float):
//res = 0
//for t in frange(0, Duck.max, 0.01):
//res += Duck.probability_density(a, t)
//if res >= p:
//return t
//raise ValueError
