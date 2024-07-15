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

func TrainAndTest(trainRatio float64, eta float64, epochs int, datasetPath string) {
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
		OutputFunc: outputFunc,
	}
	trainedPerceptron := learning.Train(eta, epochs, split.TrainingSamples, initialPerceptron)
	predict.Classify(split.TestingSamples, trainedPerceptron)
}

//  --

func CreateRandom() func() float64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() float64 {
		return rng.Float64()
	}
}

func outputFunc(weightedSum float64) float64 {
	if weightedSum > 0.0 {
		return 1.0
	}
	return -1.0
}
