package iriscmd

import (
	"github.com/ccarcaci/learn-ai/inputs/irisinput"
	"github.com/spf13/cobra"
)

var (
	IrisDatasetPath    string
	IrisRecognitionCmd = &cobra.Command{
		Use:   "iris-recognition",
		Short: "Iris Recognition",
		Long:  "Iris Recognition",
		Run: func(cmd *cobra.Command, args []string) {
			irisinput.ReadIrisDataset(IrisDatasetPath)
		},
	}
)

func init() {
	IrisRecognitionCmd.PersistentFlags().StringVar(&IrisDatasetPath, "iris-dataset-path", "", "Path to the Iris dataset")
}
