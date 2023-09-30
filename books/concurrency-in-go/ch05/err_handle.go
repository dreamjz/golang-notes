package ch05

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

func mainFunc() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	err := runJob("1")
	if err != nil {
		msg := "There was a unexpected issue, please report as a bug" // 未知异常，BUG
		var intermediateErr intermediateErr
		if errors.As(err, &intermediateErr) {
			msg = err.Error() // 已知异常
		}
		handleErr(1, err, msg)
	}
}

func handleErr(key int, err error, msg string) {
	log.SetPrefix(fmt.Sprintf("[logID: %d]", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, msg)
}

type MyError struct {
	Inner      error // 原始错误
	Message    string
	StackTrace string         // 调用栈
	Misc       map[string]any // 相关信息
}

func (err MyError) Error() string {
	return err.Message
}

func wrapError(err error, msg string) MyError {
	return MyError{
		Inner:      err,
		Message:    msg,
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]any),
	}
}

// 中层组件
type intermediateErr struct {
	error
}

func runJob(id string) error {
	jobPath := "bad/path"
	isExec, err := isGloballyExec(jobPath)

	if err != nil {
		return intermediateErr{wrapError(err, "bad/path not available")} // 封装底层错误
	} else if !isExec {
		return intermediateErr{wrapError(nil, "bad/path is not executable")}
	}

	return exec.Command(jobPath, "--id="+id).Run()
}

// 底层组件
type lowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, lowLevelErr{wrapError(err, err.Error())} // 封装系统错误
	}
	return info.Mode().Perm()&0100 == 0100, nil
}
