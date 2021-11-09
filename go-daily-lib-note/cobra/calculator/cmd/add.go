package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "do addition",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calculate(values, Add)
		fmt.Printf("%s=%.2f\n", strings.Join(args, "+"), result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
