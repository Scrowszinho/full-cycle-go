/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Category",
	Long:  `Desc`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := cmd.Flags().GetString("name")
		fmt.Println("category called ", category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("category called ", exists)
		cmd.Help()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Print("Pre run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Print("Post run")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Error")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
