package proba

import "math"

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

func PercentBack(a float64, percent float64) float64 {
	var result float64 = 0
	var i float64 = 0
	interval := IntervalCreate(percent)

	for idx := 0; i <= percent; i += 0.001 {
		result += 10 * ProbabilityDensity(a, interval[idx])
		idx += 1
	}
	return result / 100
}

func TimeBack(a, p float64, interval []float64) float64 {
	var res float64 = 0

	for idx := 0; idx < len(interval); idx += 1 {
		res += ProbabilityDensity(a, interval[idx])
		if res > p {
			return interval[idx]
		}
	}
	return 0
}

func DivMod(numerator, denominator float64) (quotient, remainder int) {
	quotient = int(numerator / denominator)
	remainder = int(numerator) % int(denominator)
	return
}
