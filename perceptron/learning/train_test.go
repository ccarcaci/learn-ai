package learning_test

import (
	"testing"

	"github.com/ccarcaci/learn-ai/perceptron"
	"github.com/ccarcaci/learn-ai/perceptron/learning"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	//  --  prepare
	trainingSample := []perceptron.Sample{
		{
			Inputs: []float64{0.0},
			Target: 0.0,
		},
	}
	initialWeights := []float64{0.0}
	initialThreshold := 0.0
	eta := 0.0
	epochs := 1

	outputFunc := func(weightedSum float64) float64 {
		if weightedSum > 0.0 {
			return 1.0
		}
		return -1.0
	}

	initialPerceptron := perceptron.Perceptron{
		Threshold:         initialThreshold,
		Weights:           initialWeights,
		OutputFunc:        outputFunc,
		EpochsErrorsCount: []int{0},
	}

	//  --  act
	trainedPerceptron := learning.Train(eta, epochs, trainingSample, initialPerceptron)

	//  --  check
	expectedPerceptron := perceptron.Perceptron{
		Threshold:         0.0,
		Weights:           []float64{0.0},
		EpochsErrorsCount: []int{0},
	}
	perceptronEqual(t, expectedPerceptron, trainedPerceptron)
}

func TestOnes(t *testing.T) {
	//  --  prepare
	trainingSample := []perceptron.Sample{
		{
			Inputs: []float64{1.0, 1.0},
			Target: 1.0,
		},
	}
	initialWeights := []float64{0.0, 0.0}
	initialThreshold := 0.5
	eta := 0.1
	epochs := 2

	outputFunc := func(weightedSum float64) float64 {
		if weightedSum > 0.0 {
			return 1.0
		}
		return -1.0
	}

	initialPerceptron := perceptron.Perceptron{
		Threshold:         initialThreshold,
		Weights:           initialWeights,
		OutputFunc:        outputFunc,
		EpochsErrorsCount: []int{0},
	}

	//  --  act
	trainedPerceptron := learning.Train(eta, epochs, trainingSample, initialPerceptron)

	//  --  check
	expectedPerceptron := perceptron.Perceptron{
		Threshold:         0.5,
		Weights:           []float64{0.0, 0.0},
		EpochsErrorsCount: []int{0, 0},
	}
	perceptronEqual(t, expectedPerceptron, trainedPerceptron)
}

func TestOnesConvergence(t *testing.T) {
	//  --  prepare
	trainingSample := []perceptron.Sample{
		{
			Inputs: []float64{1.0, 1.0},
			Target: 1.0,
		},
	}
	initialWeights := []float64{0.0, 0.0}
	initialThreshold := -0.5
	eta := 0.1
	epochs := 2

	outputFunc := func(weightedSum float64) float64 {
		if weightedSum > 0.0 {
			return 1.0
		}
		return -1.0
	}

	initialPerceptron := perceptron.Perceptron{
		Threshold:         initialThreshold,
		Weights:           initialWeights,
		OutputFunc:        outputFunc,
		EpochsErrorsCount: []int{0},
	}

	//  --  act
	trainedPerceptron := learning.Train(eta, epochs, trainingSample, initialPerceptron)

	//  --  check
	expectedPerceptron := perceptron.Perceptron{
		Threshold:         -0.3,
		Weights:           []float64{0.2, 0.2},
		EpochsErrorsCount: []int{1, 0},
	}
	perceptronEqual(t, expectedPerceptron, trainedPerceptron)
}

//  --

func perceptronEqual(t *testing.T, expected perceptron.Perceptron, trained perceptron.Perceptron) {
	assert.Equal(t, expected.Threshold, trained.Threshold)
	assert.Equal(t, expected.Weights, trained.Weights)
	assert.Equal(t, expected.EpochsErrorsCount, trained.EpochsErrorsCount)
}
