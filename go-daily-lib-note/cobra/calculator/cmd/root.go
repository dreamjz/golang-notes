package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type ErrorHandling int

const (
	ContinueOnParseError ErrorHandling = iota
	ExitOnParseError
	PanicOnParseError
	ReturnOnDividedByZero
	PanicOnDividedByZero
)

type Operation int

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

var UnrecognizedCommand = errors.New("unrecognized command")

var parseHandling int

var rootCmd = &cobra.Command{
	Use:   "Calculator",
	Short: "Simple calculator in cobra",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, UnrecognizedCommand)
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&parseHandling, "parse-error", "p", int(ContinueOnParseError), "define how command behaves if the parse fails")
}

func Execute() {
	rootCmd.Execute()
}
