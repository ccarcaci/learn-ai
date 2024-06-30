package main

import (
	"log"

	"github.com/ccarcaci/learn-ai/cmd"
)

func main() {
	if err :=cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("cmd.RootCmd.Execute(): %v", err)
	}
}
