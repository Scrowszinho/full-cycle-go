/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"teste/cli/internal/database"

	"github.com/spf13/cobra"
)

func newCreateCommand(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:  "create",
		RunE: runCreate(GetCategoryDB(GetDB())),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDB.Create(name, description)
		if err != nil {
			panic(err)
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCommand(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of category")
	createCmd.Flags().StringP("description", "d", "", "Description of category")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
