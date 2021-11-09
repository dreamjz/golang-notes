package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "do subtraction",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calculate(values, Subtract)
		fmt.Printf("%s=%.2f", strings.Join(args, "-"), result)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
