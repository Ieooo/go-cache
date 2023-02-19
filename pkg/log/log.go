package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var (
	debugLog = log.New(os.Stdout, "\033[34m[DEBUG]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[INFO ]\033[0m ", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(os.Stdout, "\033[33m[WARN ]\033[0m ", log.LstdFlags|log.Lshortfile)
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = map[LogLevel]*log.Logger{
		ErrorLevel: errorLog,
		InfoLevel:  infoLog,
		DebugLevel: debugLog,
		WarnLevel:  warnLog,
	}
	mu sync.Mutex
)

// log methods
var (
	Errorln = errorLog.Println
	Errorf  = errorLog.Printf
	Warnln  = warnLog.Println
	Warnf   = warnLog.Printf
	Infoln  = infoLog.Println
	Infof   = infoLog.Printf
	Debugln = debugLog.Println
	Debugf  = debugLog.Println
)

func SetLevel(level LogLevel) {
	mu.Lock()
	defer mu.Unlock()
	for k, v := range loggers {
		if level > k {
			v.SetOutput(ioutil.Discard)
		} else {
			v.SetOutput(os.Stdout)
		}
	}
}
