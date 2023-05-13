package chapter_13

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execCmd() {
	ls := exec.Command("ls")
	o, err := ls.CombinedOutput()
	check(err)
	fmt.Printf("Output:\n%s", string(o))

	fmt.Println(strings.Repeat("-", 20))
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	_, err = os.StartProcess(`D:\GCC\bin\ls.exe`, []string{}, procAttr)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
