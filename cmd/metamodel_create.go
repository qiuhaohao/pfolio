/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// metamodelCreateCmd represents the create command
var metamodelCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a metamodel",
	Long:  `Create a metamodel.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	metamodelCmd.AddCommand(metamodelCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metamodelCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metamodelCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
