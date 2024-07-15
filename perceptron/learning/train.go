package learning

import "github.com/ccarcaci/learn-ai/perceptron"

func Train(eta float64, epochs int, samples perceptron.Samples, initialPerceptron perceptron.Perceptron) perceptron.Perceptron {
	samplesLen := len(samples)
	trainingPerceptron := initialPerceptron

	epochsErrors := make([]int, 0)
	for i := 0; i < epochs; i++ {
		errors := 0
		for j := 0; j < samplesLen; j++ {
			sample := samples[j]
			output := Predict(sample.Inputs, trainingPerceptron)
			target := sample.Target
			update := eta * (target - output)
			delta := scalarVectProduct(update, sample.Inputs)
			trainingPerceptron.Weights = vectorsSum(trainingPerceptron.Weights, delta)
			trainingPerceptron.Threshold += update
			errors += learningErrorsCount(update)
		}
		epochsErrors = append(epochsErrors, errors)
	}
	trainingPerceptron.EpochsErrorsCount = epochsErrors
	return trainingPerceptron
}

func Predict(inputs []float64, perceptron perceptron.Perceptron) float64 {
	output := 0.0
	for i, weight := range perceptron.Weights {
		output += inputs[i] * weight
	}
	output += perceptron.Threshold
	return perceptron.OutputFunc(output)
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
