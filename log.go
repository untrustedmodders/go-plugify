package plugify

/*
#include "plugify.h"
*/
import "C"

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

func Log(msg string, sev Severity, skip int) {
	line, file, funk := caller(skip)

	C.Plugify_Log(
		msg,
		C.Severity(sev),
		C.ptrdiff_t(line),
		file,
		funk,
		plugin.name,
	)
}

var isProfiling, isLogging bool

func Scope(name string, skip int) func() {
	if !isProfiling && !isLogging {
		return func() {}
	}

	line, file, funk := caller(skip)

	var handle C.ZoneHandle

	if isProfiling {
		handle = C.Plugify_BeginZone(name, C.ptrdiff_t(line), file, funk)
	}

	if isLogging {
		C.Plugify_Log(name, C.Severity(Trace), C.ptrdiff_t(line), file, funk, plugin.name)
	}

	return func() {
		if handle != 0 {
			C.Plugify_EndZone(handle)
		}
	}
}
