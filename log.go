package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"path/filepath"
	"runtime"
	"runtime/debug"
)

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

func Log(msg string, sev Severity, info *debug.BuildInfo, skip int) {
	line, file, funk := caller(skip)

	C.Plugify_Log(
		msg,
		C.Severity(sev),
		C.ptrdiff_t(line),
		file,
		funk,
		info.Main.Path,
	)
}

var isProfiling, isLogging bool

func Scope(name string, info *debug.BuildInfo, skip int) func() {
	if !isProfiling && !isLogging {
		return func() {}
	}

	line, file, funk := caller(skip)

	var handle C.ZoneHandle

	if isProfiling {
		handle = C.Plugify_BeginZone(name, C.ptrdiff_t(line), file, funk)
	}

	if isLogging {
		C.Plugify_Log(name, C.Severity(Trace), C.ptrdiff_t(line), file, funk, info.Main.Path)
	}

	return func() {
		if handle != 0 {
			C.Plugify_EndZone(handle)
		}
	}
}

func caller(skip int) (line int, file string, funk string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	return line, filepath.Base(file), runtime.FuncForPC(pc).Name()
}
