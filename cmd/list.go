/*
Copyright © 2020 gomachan46 <shiro.gomachan46@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gomachan46/satistuffed/data"
	"github.com/spf13/cobra"
)

// listCmd represents the get command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "item list",
	Run: func(cmd *cobra.Command, args []string) {
		d := data.Load()
		itemNames := d.GetItemNames()
		for _, name := range itemNames {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//getCmd.PersistentFlags().IntVarP(&depth, "depth", "d", 10, "depth")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
