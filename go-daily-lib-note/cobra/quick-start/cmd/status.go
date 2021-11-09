package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show status of the git repository",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "status", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
