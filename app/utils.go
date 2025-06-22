package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func (app *application) errorTrace(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())	
	app.errLog.Output(2, trace)
	os.Exit(1)
}
