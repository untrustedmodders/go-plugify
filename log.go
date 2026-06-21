package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

var isProfiling, isLogging bool

type Severity int

const (
	Unknown Severity = iota
	Trace
	Debug
	Info
	Warning
	Error
	Fatal
)

func Log(msg string, sev Severity, mod string, skip int) {
	line, file, funk := caller(skip)

	C.Plugify_Log(
		msg,
		C.Severity(sev),
		C.ptrdiff_t(line),
		file,
		funk,
		mod,
	)
}

func Scope(name string, mod string, skip int) func() {
	if !isProfiling && !isLogging {
		return func() {}
	}

	line, file, funk := caller(skip)

	var handle C.ZoneHandle

	if isProfiling {
		handle = C.Plugify_BeginZone(name, C.ptrdiff_t(line), file, funk)
	}

	if isLogging {
		C.Plugify_Log(name, C.Severity(Trace), C.ptrdiff_t(line), file, funk, mod)
	}

	return func() {
		if handle != 0 {
			C.Plugify_EndZone(handle)
		}
	}
}

func stacktrace(err any) {
	msg := fmt.Sprintf("%v", err)
	stack := debug.Stack()
	if len(stack) > 0 {
		msg += fmt.Sprintf("\nStack Trace: \n%s", stack)
	}
	Log(msg, Error, plg().name, 3)
}

func caller(skip int) (line int, file string, funk string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	return line, filepath.Base(file), runtime.FuncForPC(pc).Name()
}

func panicker(err any) {
	stacktrace(err)
	panic(err)
}
