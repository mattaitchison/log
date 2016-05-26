package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var std = New(os.Stderr, "", log.Lmicroseconds)

type Logger struct {
	*log.Logger
	fatal *log.Logger
	debug *log.Logger
}

func New(out io.Writer, prefix string, flag int) *Logger {
	if prefix == "" {
		prefix = " "
	}
	logger := &Logger{
		Logger: log.New(out, "[INFO]"+prefix, flag),
		fatal:  log.New(out, "[FATAL]"+prefix, flag|log.Lshortfile),
		debug:  log.New(ioutil.Discard, "[DEBUG]"+prefix, flag|log.Lshortfile),
	}

	if os.Getenv("DEBUG") != "" {
		logger.SetDebug()
	}
	return logger
}

func (l *Logger) SetDebug() {
	l.SetFlags(l.Flags() | log.Lshortfile)
	l.SetDebugOutput(os.Stderr)
}

func (l *Logger) SetDebugOutput(w io.Writer) {
	l.debug.SetOutput(w)
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.debug.Output(2, fmt.Sprintln(v...))
}

func SetDebug() { std.SetDebug() }

func SetDebugOutput(w io.Writer) { std.SetDebugOutput(w) }

func Debug(v ...interface{}) { std.debug.Output(2, fmt.Sprint(v...)) }

func Debugf(format string, v ...interface{}) { std.debug.Output(2, fmt.Sprintf(format, v...)) }

func Debugln(v ...interface{}) { std.debug.Output(2, fmt.Sprintln(v...)) }

func Fatal(v ...interface{}) {
	std.fatal.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	std.fatal.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	std.fatal.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
