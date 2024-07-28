package iriscmd

import (
	"github.com/ccarcaci/learn-ai/models/irismodel"
	"github.com/spf13/cobra"
)

var (
	IrisDatasetPath    string
	TrainingType       string
	ActivationType     string
	TrainingRatio      float64
	Eta                float64
	Epochs             int
	IrisRecognitionCmd = &cobra.Command{
		Use:   "iris-recognition",
		Short: "Iris Recognition",
		Long:  "Iris Recognition",
		Run: func(cmd *cobra.Command, args []string) {
			irismodel.TrainAndTest(TrainingRatio, Eta, Epochs, IrisDatasetPath, TrainingType, ActivationType)
		},
	}
)

func init() {
	IrisRecognitionCmd.PersistentFlags().StringVar(&IrisDatasetPath, "iris-dataset-path", "", "Path to the Iris dataset")
	IrisRecognitionCmd.PersistentFlags().StringVar(&TrainingType, "training-type", "", "Training type batch/online")
	IrisRecognitionCmd.PersistentFlags().StringVar(&ActivationType, "activation-type", "", "Activation type adaline/perceptron")
	IrisRecognitionCmd.PersistentFlags().Float64Var(&Eta, "eta", 0.1, "Learning rate")
	IrisRecognitionCmd.PersistentFlags().IntVar(&Epochs, "epochs", 100, "Number of epochs")
	IrisRecognitionCmd.PersistentFlags().Float64Var(&TrainingRatio, "training-ratio", 0.8, "Training ratio")
}
