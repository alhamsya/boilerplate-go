package customLog

import (
	"runtime"
	"strings"
)

func trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()[strings.LastIndex(runtime.FuncForPC(pc).Name(), "/")+1:]
}