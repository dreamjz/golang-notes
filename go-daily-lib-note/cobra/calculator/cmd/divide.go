package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var divideByZeroHandling int

var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "do division",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calculate(values, Divide)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "/"), result)
	},
}

func init() {
	divideCmd.Flags().IntVarP(&divideByZeroHandling, "divided-by-zero", "d", int(PanicOnDividedByZero), "define divide command behaves if divided by zero")
	rootCmd.AddCommand(divideCmd)
}
