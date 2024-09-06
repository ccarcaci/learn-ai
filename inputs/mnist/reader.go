package mnist

import (
	"encoding/binary"
	"fmt"
	"os"
)

type MNISTEntry struct {
	Label uint8
	Image []uint8
}

const labelsMagicNumber = 0x00000801
const datasetMagicNumber = 0x00000803

func ReadMNISTDataset(labelsPath string, datasetPath string) ([]MNISTEntry, error) {
	labelsFile, err := os.Open(labelsPath)
	if err != nil {
		return nil, fmt.Errorf("labels os.Open(): %w", err)
	}
	defer labelsFile.Close()

	datasetFile, err := os.Open(datasetPath)
	if err != nil {
		return nil, fmt.Errorf("dataset os.Open(): %w", err)
	}
	defer datasetFile.Close()

	err = expectMagicNumber(labelsFile, labelsMagicNumber)
	if err != nil {
		return nil, fmt.Errorf("expectMagicNumber(): %d %w", labelsMagicNumber, err)
	}

	err = expectMagicNumber(datasetFile, datasetMagicNumber)
	if err != nil {
		return nil, fmt.Errorf("expectMagicNumber(): %d %w", datasetMagicNumber, err)
	}


	return []MNISTEntry{}, nil
//	labelsCount := int32(0)
//	err = getItemsCount(file, &itemsCount)
//	
//	mnistEntries := make([]MNISTEntry, itemsCount)
//	err = readMNISTEntries(file, mnistEntries)
}

//  --

func expectMagicNumber(file *os.File, expected int32) error {
	magicNumber := int32(0)
	err := binary.Read(file, binary.BigEndian, &magicNumber)
	if err != nil {
		return fmt.Errorf("binary.Read(): %w", err)
	}
	if magicNumber != expected {
		return fmt.Errorf("invalid magic number: %d", magicNumber)
	}
	return nil
}

func getItemsCount(file *os.File, itemsCount *int32) error {
	err := binary.Read(file, binary.BigEndian, itemsCount)
	if err != nil {
		return fmt.Errorf("binary.Read(): %w", err)
	}
	return nil
}
