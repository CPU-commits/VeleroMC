package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var engineCommand = &cobra.Command{
	Use:   "engine",
	Short: "Query engine information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Subcommands
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all engines",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available engines")
		// Find
		engines, err := enginesRepository.FindAll()
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, engine := range engines {
			fmt.Printf("%d. %s\n", i+1, engine.Name)
		}
	},
}

var versionsCmd = &cobra.Command{
	Use:   "versions [engine_name]",
	Short: "List versions of a specific engine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		engineName := args[0]
		// Find
		engine, err := enginesRepository.FindOneByEngine(engineName)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, version := range engine.Versions {
			fmt.Printf("%d. %s\n", i+1, version)
		}
	},
}

func init() {
	engineCommand.AddCommand(lsCmd)
	engineCommand.AddCommand(versionsCmd)

	rootCmd.AddCommand(engineCommand)
}
