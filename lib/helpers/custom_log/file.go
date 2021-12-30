package customLog

import (
	"os"
	"path/filepath"
)

func OpenFile(logFile string) (*os.File, error) {
	if logFile == "" {
		return nil, nil
	}

	err := os.MkdirAll(filepath.Dir(logFile), 0755)
	if err != nil && err != os.ErrExist {
		return nil, err
	}

	return os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}