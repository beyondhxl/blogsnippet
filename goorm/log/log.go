package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorlog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infolog  = log.New(os.Stdout, "\033[34m[info]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorlog, infolog}
	mu       sync.Mutex
)

var (
	Error  = errorlog.Println
	Errorf = errorlog.Printf
	Info   = infolog.Println
	Infof  = infolog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorlog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infolog.SetOutput(ioutil.Discard)
	}
}
