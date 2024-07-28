package perceptron

type Sample struct {
	Inputs []float64
	Target float64
}

type Samples []Sample

type Perceptron struct {
	Threshold  float64
	Weights    []float64
	OutputFunc func(float64) float64
}

//  --

func Linear(x float64) float64 {
	return x
}

func Step(x float64) float64 {
	if x > 0.0 {
		return 1.0
	}
	return -1.0
}

//  --

func BatchErr(threshold float64, inputs []float64, weights []float64, activationFunc func(float64) float64) float64 {
	output := 0.0
	for i, weight := range weights {
		output += inputs[i] * weight
	}
	output += threshold
	return activationFunc(output)
}

func OnlineErr(threshold float64, input float64, weight float64, activationFunc func(float64) float64) float64 {
	return activationFunc(weight*input + threshold)
}
