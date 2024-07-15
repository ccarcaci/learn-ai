package irismodel

import (
	"github.com/ccarcaci/learn-ai/inputs/irisinput"
	"github.com/ccarcaci/learn-ai/perceptron"
)

func MapSamples(irisData []irisinput.IrisData) perceptron.Samples {
	samples := make([]perceptron.Sample, 0)
	for _, irisData := range irisData {
		samples = append(samples, perceptron.Sample{
			Inputs: []float64{irisData.SepalLength, irisData.PetalLength},
			Target: irisData.IrisType,
		})
	}
	return samples
}
