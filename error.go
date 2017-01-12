package main

import (
	"fmt"
	"os"
	"runtime"

	ansi "github.com/k0kubun/go-ansi"
)

func assert(ok bool, s string) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		ansi.Printf(fmt.Sprintf("\n%s:%d\n[error] \x1b[31m%s\x1b[0m\n", file, line, s))
		os.Exit(1)
	}
}
