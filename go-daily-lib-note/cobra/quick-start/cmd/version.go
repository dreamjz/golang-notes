package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show git version info",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprintf(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
