package pkg

import (
	"fmt"
	"runtime"
	"strings"
)

func CaptureStackTrace() string {
	var sb strings.Builder
	pcs := make([]uintptr, 1)
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s:%d", frame.File, frame.Line))
		if !more {
			break
		}
	}

	return sb.String()
}
