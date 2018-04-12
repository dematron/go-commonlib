package commonlib

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/go-errors/errors"
)

// CheckError handle incoming errors and support additional options
// such as trace and logging trace to file
func CheckError(inerr error, errorTrace bool, errorTracepath string) error {
	if inerr != nil {
		log.Printf("Error: %v%s", inerr, Carriage)
		// trace
		pc := make([]uintptr, 15)
		n := runtime.Callers(3, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		log.Printf("%s,:%d %s%s", frame.File, frame.Line, frame.Function, Carriage)
		if errorTrace {
			traceFile, err := os.OpenFile(errorTracepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println(err)
			}
			defer traceFile.Close()
			t := time.Now().Format("2006/01/02 15:04:05 MST")
			e := errors.New(inerr)
			traceFile.WriteString(t + Carriage)
			traceFile.WriteString(e.TypeName() + Carriage)
			traceFile.Write(e.Stack())
			// Issue a Sync to flush writes to stable storage.
			traceFile.Sync()
		}
		//
	}
	return inerr
}
