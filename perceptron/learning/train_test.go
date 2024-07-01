package learning_test

import (
	"testing"

	"github.com/ccarcaci/learn-ai/perceptron/learning"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	//  --  prepare
	trainingSample := []learning.TrainSingleSample{
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

	//  --  act
	trainedPerceptron := learning.Learn(eta, epochs, trainingSample, initialThreshold, initialWeights, outputFunc)

	//  --  check
	expectedPerceptron := learning.Perceptron{
		Threshold:         0.0,
		Weights:           []float64{0.0},
		EpochsErrorsCount: []int{0},
	}
	assert.Equal(t, expectedPerceptron, trainedPerceptron)
}

func TestOnes(t *testing.T) {
	//  --  prepare
	trainingSample := []learning.TrainSingleSample{
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

	//  --  act
	trainedPerceptron := learning.Learn(eta, epochs, trainingSample, initialThreshold, initialWeights, outputFunc)

	//  --  check
	expectedPerceptron := learning.Perceptron{
		Threshold:         0.5,
		Weights:           []float64{0.0, 0.0},
		EpochsErrorsCount: []int{0, 0},
	}
	assert.Equal(t, expectedPerceptron, trainedPerceptron)
}

func TestOnesConvergence(t *testing.T) {
	//  --  prepare
	trainingSample := []learning.TrainSingleSample{
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

	//  --  act
	trainedPerceptron := learning.Learn(eta, epochs, trainingSample, initialThreshold, initialWeights, outputFunc)

	//  --  check
	expectedPerceptron := learning.Perceptron{
		Threshold:         -0.3,
		Weights:           []float64{0.2, 0.2},
		EpochsErrorsCount: []int{1, 0},
	}
	assert.Equal(t, expectedPerceptron, trainedPerceptron)
}
