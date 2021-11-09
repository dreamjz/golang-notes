package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sim-git",
	Short: "Sim-git is a simulation of git",
	Long: `Sim-git is a simulation of git.
Git is a free and open source distributed version control system-designed to 
handle everything from small to very large projects with speed and efficiency.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}
