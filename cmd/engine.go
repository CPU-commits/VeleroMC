package cmd

import (
	"fmt"

	"github.com/CPU-commits/VeleroMC/engine"
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
		engine := engine.EngineFactory(engineName)
		engineVersions, err := engine.List()
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, version := range engineVersions {
			fmt.Printf("%d. %s\n", i+1, version)
		}
	},
}

var buildCmd = &cobra.Command{
	Use:   "build [engine_name] [version]",
	Short: "Download and build minecraft server",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		engineName := args[0]
		version := args[1]
		// Download java server file
		engine := engine.EngineFactory(engineName)
		file, err := engine.Download(version)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Build minecraft server
		err = engine.Build(file)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	},
}

var runCmd = &cobra.Command{
	Use:   "run [engine_name] [path]",
	Short: "Run a builded minecraft server",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		engineName := args[0]
		serverPath := args[1]
		// Run
		engine := engine.EngineFactory(engineName)
		if err := engine.Run(serverPath); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	},
}

func init() {
	engineCommand.AddCommand(lsCmd)
	engineCommand.AddCommand(versionsCmd)
	engineCommand.AddCommand(buildCmd)
	engineCommand.AddCommand(runCmd)

	rootCmd.AddCommand(engineCommand)
}
