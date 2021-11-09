package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var DividedByZero = errors.New("divided by zero")

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%s", cmd.Name(), args, err)
	os.Exit(1)
}

func ConvertArgsToFloat64Slice(args []string, errorHandling ErrorHandling) []float64 {
	result := make([]float64, 0, len(args))
	for _, arg := range args {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			switch errorHandling {
			case ExitOnParseError:
				fmt.Fprintf(os.Stderr, "invalid number: %s \n", arg)
				os.Exit(1)
			case PanicOnParseError:
				panic(err)
			}
		}
		result = append(result, val)
	}
	return result
}

func calculate(values []float64, operation Operation) float64 {
	var result float64
	if len(values) == 0 {
		return result
	}
	result = values[0]
	for i := 1; i < len(values); i++ {
		switch operation {
		case Add:
			result += values[i]
		case Subtract:
			result -= values[i]
		case Multiply:
			result *= values[i]
		case Divide:
			if values[i] == 0 {
				switch ErrorHandling(divideByZeroHandling) {
				case ReturnOnDividedByZero:
					return result
				case PanicOnDividedByZero:
					panic(DividedByZero)
				}
			}
			result /= values[i]
		}
	}
	return result
}
