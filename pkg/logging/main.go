package logging

import (
	f "fmt"
	"os"
)

func LogErrorExitf(msg string, args ...interface{}) {
        f.Fprintf(os.Stderr, msg+"\n", args...)
        os.Exit(1)
}

func LogErrorf(msg string, args ...interface{}) {
        f.Fprintf(os.Stderr, msg+"\n", args...)
}
