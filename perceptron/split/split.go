package split

import "github.com/ccarcaci/learn-ai/perceptron"

type GenerateRandom func() float64

const (
	Training SetName = "training"
	Testing  SetName = "testing"
)

type SetName string

type SamplesSplit struct {
	TrainingSamples perceptron.Samples
	TestingSamples  perceptron.Samples
}

func SplitSamples(samples perceptron.Samples, splitRatio float64, randomGenerator GenerateRandom) SamplesSplit {
	trainingSamples := make([]perceptron.Sample, 0)
	testingSamples := make([]perceptron.Sample, 0)

	for _, sample := range samples {
		ratio := getRatio(len(trainingSamples), len(testingSamples))
		if GetChoice(ratio, splitRatio, randomGenerator) == Training {
			trainingSamples = append(trainingSamples, sample)
		} else {
			testingSamples = append(testingSamples, sample)
		}
	}

	return SamplesSplit{
		TrainingSamples: trainingSamples,
		TestingSamples:  testingSamples,
	}
}

//  --

func GetChoice(ratio float64, splitRatio float64, randomGenerator GenerateRandom) SetName {
	if ratio > splitRatio {
		return confirmTesting(randomGenerator)
	}
	return confirmTraining(randomGenerator)
}

func confirmTesting(randomGenerator GenerateRandom) SetName {
	if randomGenerator() >= 0.9 {
		return Training
	}
	return Testing
}

func confirmTraining(randomGenerator GenerateRandom) SetName {
	if randomGenerator() >= 0.9 {
		return Testing
	}
	return Training
}

func getRatio(trainingLen int, testingLen int) float64 {
	if testingLen + trainingLen == 0 {
		return 1.0
	}
	return float64(trainingLen) / float64(trainingLen + testingLen)
}
