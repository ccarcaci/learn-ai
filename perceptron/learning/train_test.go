package learning_test

import (
	"math"
	"testing"

	"github.com/ccarcaci/learn-ai/perceptron/learning"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	//  --  prepare
	input := []float64{0.0}
	weights := []float64{0.0}
	target := 0.0
	eta := 0.0
	epochs := 1

	//  --  act
	trainedWeights := learning.Learn(input, weights, target, eta, epochs)

	//  --  check
	expectedWeights := []float64{0.0}
	assert.InDelta(t, expectedWeights[0], trainedWeights[0], math.SmallestNonzeroFloat64)
}

func TestOnes(t *testing.T) {
	//  --  prepare
	input := []float64{1.0, 1.0}
	weights := []float64{0.0, 0.0}
	target := 1.0
	eta := 0.1
	epochs := 2

	//  --  act
	trainedWeights := learning.Learn(input, weights, target, eta, epochs)

	//  -- check
	expectedWeigths := []float64{0.18, 0.18}
	assertFloatArrayEqual(t, expectedWeigths, trainedWeights)
}

//  --

const minEpsilon = 2.7755575615628914e-17

func assertFloatArrayEqual(t *testing.T, expected []float64, actual []float64) bool {
	t.Helper()

	lenExpected := len(expected)
	lenActual := len(actual)
	if len(expected) != len(actual) {
		return assert.Fail(t, "expected len is %d, actual len is %d", lenExpected, lenActual)
	}
	for i := 0; i < len(expected); i++ {
		assert.InDelta(t, expected[i], actual[i], minEpsilon)
	}
	return true
}
