package learning

type TrainSingleSample struct {
	Inputs []float64
	Target float64
}

type TrainSamples []TrainSingleSample

type OutputFunc func(float64) float64

type Perceptron struct {
	Threshold         float64
	Weights           []float64
	EpochsErrorsCount []int
}

func Learn(eta float64, epochs int, samples TrainSamples, initialThreshold float64, initialWeights []float64, targets []float64, outputFunc OutputFunc) Perceptron {
	samplesLen := len(samples)
	weights := initialWeights
	threshold := initialThreshold
	epochsErrors := make([]int, 0)
	for i := 0; i < epochs; i++ {
		errors := 0
		for j := 0; j < samplesLen; j++ {
			output := Predict(samples[j].Inputs, weights, threshold, outputFunc)
			target := targets[j]
			update := eta * (target - output)
			delta := scalarVectProduct(update, samples[j].Inputs)
			weights = vectorsSum(weights, delta)
			threshold += update
			errors += learningErrorsCount(update)
		}
		epochsErrors = append(epochsErrors, errors)
	}
	return Perceptron{threshold, weights, epochsErrors}
}

func Predict(inputs []float64, weights []float64, threshold float64, outputFunc OutputFunc) float64 {
	output := 0.0
	for i, weight := range weights {
		output += inputs[i] * weight
	}
	return outputFunc(output)
}

//  --

func scalarVectProduct(scalar float64, vector []float64) []float64 {
	vectLen := len(vector)
	result := make([]float64, vectLen)
	for i := 0; i < vectLen; i++ {
		result[i] = scalar * vector[i]
	}
	return result
}

func vectorsSum(first []float64, second []float64) []float64 {
	vectLen := len(first)
	result := make([]float64, vectLen)
	for i := 0; i < vectLen; i++ {
		result[i] = first[i] + second[i]
	}
	return result
}

func learningErrorsCount(update float64) int {
	if update != 0 {
		return 1
	}
	return 0
}
