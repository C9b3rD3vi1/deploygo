/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Manage environments",
	Long: `Manage environments for your application.
This command allows you to create, update, and delete environments for your application.
You can use this command to manage your application's environments, such as development, staging, and production.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("environment called")
	},
}

func init() {
	rootCmd.AddCommand(environmentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environmentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
