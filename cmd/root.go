package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Satistuffed",
	Short: "Calculate 100% efficient (stuffed) for Satisfactory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Satistuffed!")
	},
}
