package irismodel_test

import (
	"testing"

	"github.com/ccarcaci/learn-ai/inputs/irisinput"
	"github.com/ccarcaci/learn-ai/models/irismodel"
	"github.com/ccarcaci/learn-ai/perceptron"
	"github.com/stretchr/testify/assert"
)

func TestMapSamples(t *testing.T) {
	//  --  prepare
	inputIrisData := []irisinput.IrisData{
		{
			SepalLength: 0.0,
			PetalLength: 1.0,
			IrisType:    -1.0,
		},
		{
			SepalLength: 1.0,
			PetalLength: 2.0,
			IrisType:    1.0,
		},
	}

	//  --  act
	samples := irismodel.MapSamples(inputIrisData)

	//  --  check
	assert.Equal(t, 2, len(samples))
	expected := perceptron.Samples{
		{
			Inputs: []float64{0.0, 1.0},
			Target: -1.0,
		},
		{
			Inputs: []float64{1.0, 2.0},
			Target: 1.0,
		},
	}
	assert.Equal(t, expected, samples)
}
