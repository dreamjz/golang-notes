package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "do multiplication",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calculate(values, Multiply)
		fmt.Printf("%s=%.2f\n", strings.Join(args, "*"), result)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
}
