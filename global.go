package plugify

/*
#include "plugify.h"
*/
import "C"

var baseDir = ""
var extensionsDir = ""
var configsDir = ""
var dataDir = ""
var logsDir = ""
var cacheDir = ""

func BaseDir() string {
	return baseDir
}

func ExtensionsDir() string {
	return extensionsDir
}

func ConfigsDir() string {
	return configsDir
}

func DataDir() string {
	return dataDir
}

func LogsDir() string {
	return logsDir
}

func CacheDir() string {
	return cacheDir
}

func IsLoaded(name string, constraint string) bool {
	return bool(C.Plugify_IsLoaded(name, constraint))
}
