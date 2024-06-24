package learning

func Learn(input []float64, weights []float64, target float64, eta float64, epochs int) []float64 {
	weightsLen := len(weights)
	for i := 0; i < epochs; i++ {
		output := Predict(input, weights)
		for j := 0; j < weightsLen; j++ {
			delta := eta * (target - output) * input[j]
			weights[j] += delta
		}
	}
	return weights
}

func Predict(input []float64, weights []float64) float64 {
	output := 0.0
	for i, weight := range weights {
		output += input[i] * weight
	}
	return output
}
