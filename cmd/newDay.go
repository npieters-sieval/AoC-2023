/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	day string
)

// newDayCmd represents the newDay command
var newDayCmd = &cobra.Command{
	Use:   "newDay",
	Short: "A command that adds a new day to the project",
	Long: `The command adds a new day with the follwing items:
		* Tests
		* InputFile
		* TestFile
`,
	Run: func(cmd *cobra.Command, args []string) {
		folderName := "Day-" + day
		if err := os.Mkdir(folderName, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newDayCmd)

	newDayCmd.Flags().StringVarP(&day, "day", "d", "0", "Specify a day")
}
