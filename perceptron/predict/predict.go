package predict

import (
	"log"

	"github.com/ccarcaci/learn-ai/perceptron"
)

func Classify(samples perceptron.Samples, perceptron perceptron.Perceptron) {
	for _, sample := range samples {
		out := Predict(sample.Inputs, perceptron)
		log.Println("inputs: ", sample.Inputs, " | target: ", sample.Target, " | output: ", out)
	}
}

func Predict(inputs []float64, perceptron perceptron.Perceptron) float64 {
	output := 0.0
	for i, weight := range perceptron.Weights {
		output += inputs[i] * weight
	}
	output += perceptron.Threshold
	return perceptron.OutputFunc(output)
}
