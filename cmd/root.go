package cmd

import (
	"log"
	"os"

	"github.com/ccarcaci/learn-ai/cmd/iris"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "leai",
	Short: "Leai a powerful command line AI tool",
	Long:  "Leai is a command line that provides a set of examples and AI tools",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		cmd.Help()
	//	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("rootCmd.Execute(): %v", err)
	}
	os.Exit(0)
}

func init() {
	RootCmd.AddCommand(iris.IrisRecognitionCmd)
}
