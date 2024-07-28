package learning

import "github.com/ccarcaci/learn-ai/perceptron"

func Train(eta float64, epochs int, trainingType string, samples perceptron.Samples, initialPerceptron perceptron.Perceptron, activationFunc func(float64) float64) perceptron.Perceptron {
	trainingPerceptron := initialPerceptron

	for i := 0; i < epochs; i++ {
		if trainingType == "online" {
			trainingPerceptron = onlineLearning(eta, samples, trainingPerceptron, activationFunc)
			continue
		}
		if trainingType == "batch" {
			trainingPerceptron = batchLearning(eta, samples, trainingPerceptron, activationFunc)
		}
	}
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

func batchLearning(eta float64, samples perceptron.Samples, perceptron perceptron.Perceptron, activationFunc func(float64) float64) perceptron.Perceptron {
	partials := make([]float64, len(perceptron.Weights))
	thresholdErrFeed := 0.0
	for _, sample := range samples {
		sum := perceptron.Threshold
		for i, input := range sample.Inputs {
			sum += input * perceptron.Weights[i]
		}

		activation := activationFunc(sum)
		errFeed := sample.Target - activation
		thresholdErrFeed += errFeed

		for i, input := range sample.Inputs {
			partials[i] = errFeed * input
		}
	}

	for i, partial := range partials {
		perceptron.Weights[i] += eta * partial
	}
	perceptron.Threshold += eta * thresholdErrFeed

	return perceptron
}

func onlineLearning(eta float64, samples perceptron.Samples, perceptron perceptron.Perceptron, activationFunc func(float64) float64) perceptron.Perceptron {
	for _, sample := range samples {
		sum := perceptron.Threshold
		for i, input := range sample.Inputs {
			sum += input * perceptron.Weights[i]
		}

		activation := activationFunc(sum)
		errFeed := eta * (sample.Target - activation)

		for i, input := range sample.Inputs {
			perceptron.Weights[i] += errFeed * input
		}
		perceptron.Threshold += errFeed
	}

	return perceptron
}
