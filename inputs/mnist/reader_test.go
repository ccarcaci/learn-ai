package mnist_test

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fileContent struct {
	magicNumber int32
	itemsCount  int32

	labels []uint8
}

func TestDiscoveryReadTestingLabels(t *testing.T) {
	testingLabelsPath := "../../datasets/t10k-labels-idx1-ubyte"
	file, err := os.Open(testingLabelsPath)
	defer func() {
		err := file.Close()
		if err != nil {
			os.Exit(-1)
		}
	}()
	require.NoError(t, err)

	//  --
	magicNumber := int32(0)
	err = binary.Read(file, binary.BigEndian, &magicNumber)
	require.NoError(t, err)
	assert.Equal(t, int32(0x00000801), magicNumber)
	t.Log("magic number: ", magicNumber)

	//  --
	itemsCount := int32(0)
	err = binary.Read(file, binary.BigEndian, &itemsCount)
	require.NoError(t, err)
	t.Log("items count: ", itemsCount)

	//  --
	data := make([]uint8, itemsCount)
	err = binary.Read(file, binary.BigEndian, &data)
	require.NoError(t, err)
	t.Log("labels: ", data)
}

func TestDiscoveryReadTestingImages(t *testing.T) {
	testingImagesPath := "../../datasets/t10k-images-idx3-ubyte"
	file, err := os.Open(testingImagesPath)
	defer func() {
		err := file.Close()
		if err != nil {
			os.Exit(-1)
		}
	}()
	require.NoError(t, err)

	//  --
	magicNumber := int32(0)
	err = binary.Read(file, binary.BigEndian, &magicNumber)
	require.NoError(t, err)
	assert.Equal(t, int32(0x00000803), magicNumber)
	t.Log("magic number: ", magicNumber)

	//  --
	itemsCount := int32(0)
	err = binary.Read(file, binary.BigEndian, &itemsCount)
	require.NoError(t, err)
	t.Log("items count: ", itemsCount)

	//  --
	cols := int32(0)
	err = binary.Read(file, binary.BigEndian, &cols)
	require.NoError(t, err)
	t.Log("cols: ", cols)

	//  --
	rows := int32(0)
	err = binary.Read(file, binary.BigEndian, &rows)
	require.NoError(t, err)
	t.Log("rows: ", rows)

	//  --
	data := make([]uint8, 28*28)
	err = binary.Read(file, binary.BigEndian, &data)
	require.NoError(t, err)
	printNumber(t, rows, cols, data)
}

//  --

func printNumber(t *testing.T, cols int32, rows int32, data []uint8) {
	t.Helper()
	var row string

	for i := int32(0); i < rows; i++ {
		row = ""
		for j := int32(0); j < cols; j++ {
			dot := grayScaleToAscii(data[i*cols+j])
			row = fmt.Sprintf("%s %s", row, dot)
		}
		t.Log(row)
	}
}

const ramp = "@#+=."

func grayScaleToAscii(grayScale uint8) string {
	if grayScale >= 255 {
		return "."
	}
	index := grayScale / 51
	return string(ramp[index])
}
