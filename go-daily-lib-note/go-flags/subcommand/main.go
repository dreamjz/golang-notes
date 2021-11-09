package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"strconv"
	"strings"
)

var InvalidOperation = errors.New("invalid operation")

type MathCommand struct {
	Op     string `long:"op" description:"operation to execute"`
	Args   []string
	Result int64
}

func (mc *MathCommand) Execute(args []string) error {
	op := mc.Op
	if op != "+" && op != "-" && op != "x" && op != "/" {
		return InvalidOperation
	}
	// make([]T,len[,cap])
	nums := make([]int64, 0, len(args))
	for _, arg := range args {
		num, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}
		nums = append(nums, num)
	}
	mc.Result = Calculate(nums, op)
	mc.Args = args
	return nil
}

func Calculate(nums []int64, op string) int64 {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		switch op {
		case "+":
			result += nums[i]
		case "-":
			result -= nums[i]
		case "x":
			result *= nums[i]
		case "/":
			result /= nums[i]
		}
	}
	return result
}

type Option struct {
	Math MathCommand `command:"math"`
}

func main() {
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		return
	}
	fmt.Printf("The result of %s is %d ", strings.Join(opt.Math.Args, opt.Math.Op), opt.Math.Result)
}
