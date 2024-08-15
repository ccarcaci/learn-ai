package mnist_test

import (
	"encoding/binary"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type fileContent struct {
	magicNumber int32
	itemsCount  int32

	labels []uint8
}

func TestDiscoveryReadFirst10TestingLabels(t *testing.T) {
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
