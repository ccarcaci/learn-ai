package irismodel

import (
	"log"
	"math/rand"
	"time"

	"github.com/ccarcaci/learn-ai/inputs/irisinput"
	"github.com/ccarcaci/learn-ai/perceptron"
	"github.com/ccarcaci/learn-ai/perceptron/learning"
	"github.com/ccarcaci/learn-ai/perceptron/predict"
	"github.com/ccarcaci/learn-ai/perceptron/split"
)

func TrainAndTest(trainRatio float64, eta float64, epochs int, datasetPath string, trainingType string, activationType string) {
	irisDataset, err := irisinput.ReadIrisDataset(datasetPath)
	if err != nil {
		log.Fatalf("inputs.ReadIrisDataset(): %v", err)
	}

	samples := MapSamples(irisDataset)
	split := split.SplitSamples(samples, trainRatio, CreateRandom())
	log.Printf("training set size: %d\n", len(split.TrainingSamples))
	log.Printf("testing set size: %d\n", len(split.TestingSamples))

	initialWeights := []float64{0.0, 0.0}
	initialThreshold := 0.0
	initialPerceptron := perceptron.Perceptron{
		Threshold:  initialThreshold,
		Weights:    initialWeights,
		OutputFunc: perceptron.Step,
	}

	activationFunc := perceptron.Step
	if activationType == "adaline" {
		activationFunc = perceptron.Linear
	}

	trainedPerceptron := learning.Train(eta, epochs, trainingType, split.TrainingSamples, initialPerceptron, activationFunc)
	predict.Classify(split.TestingSamples, trainedPerceptron)
}

//  --

func CreateRandom() func() float64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() float64 {
		return rng.Float64()
	}
}
