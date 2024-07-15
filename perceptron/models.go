package perceptron

type Sample struct {
	Inputs []float64
	Target float64
}

type Samples []Sample

type Perceptron struct {
	Threshold         float64
	Weights           []float64
	OutputFunc        func(float64) float64
	EpochsErrorsCount []int
}
