package irisds

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type IrisData struct {
	SepalLength float64
	PetalLength float64
	IrisType    float64
}

func ReadIrisDataset(datasetPath string) ([]IrisData, error) {
	file, err := os.Open(datasetPath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(): %w", err)
	}
	defer file.Close()

	irisDataset := make([]IrisData, 0)

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reader.Read(): %w", err)
		}
		irisData, err := recordToIrisData(record)
		if err != nil {
			return nil, fmt.Errorf("recordToIrisData(record): %w", err)
		}
		irisDataset = append(irisDataset, irisData)
	}
	log.Printf("Dataset size: %d\n", len(irisDataset))
	return irisDataset, nil
}

//  --

func recordToIrisData(record []string) (IrisData, error) {
	sepalLength, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		return IrisData{}, fmt.Errorf("strconv.ParseFloat(record[0], 64): %w", err)
	}
	petalLength, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return IrisData{}, fmt.Errorf("strconv.ParseFlow(record[2], 64): %w", err)
	}
	irisType := -1.0
	if record[4] == "Iris-setosa" {
		irisType = 1.0
	}
	return IrisData{
		SepalLength: sepalLength,
		PetalLength: petalLength,
		IrisType:    irisType,
	}, nil
}
