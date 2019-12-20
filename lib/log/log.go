package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger .
type Logger interface {
	Printf(format string, v ...interface{})
	Debug(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	ErrorIf(err error, v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})
	SetOutputFile(name string)
	SetWriter(w io.Writer)
	GetWriter() io.Writer
}

// NewLogger .
func NewLogger() Logger {
	return newLogger("global", os.Stderr)
}

// NewLoggerWithName .
func NewLoggerWithName(name string) Logger {
	return newLogger(name, os.Stderr)
}

// NewLoggerWithNameAndFile .
func NewLoggerWithNameAndFile(name string, file string) Logger {
	l := newLogger(name, os.Stderr)
	l.SetOutputFile(file)
	return l
}

var loggerManager = []*logger{}

func newLogger(name string, out io.Writer) *logger {
	l := &logger{
		impl: &logrus.Logger{
			Out:       out,
			Formatter: &serverXTextFormatter{},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
		name: name,
		mu:   &sync.Mutex{},
	}
	l.impl.AddHook(serverXContextHook{})
	loggerManager = append(loggerManager, l)
	return l
}

type logger struct {
	impl *logrus.Logger
	name string
	mu   *sync.Mutex
}

var enableDebug = true

// EnableDebug .
func EnableDebug(enable bool) {
	enableDebug = enable
	if enableDebug {
		for _, v := range loggerManager {
			v.impl.SetLevel(logrus.DebugLevel)
		}
	} else {
		for _, v := range loggerManager {
			v.impl.SetLevel(logrus.InfoLevel)
		}
	}
}

// SwitchEnableDebug .
func SwitchEnableDebug() bool {
	if enableDebug {
		EnableDebug(false)
	} else {
		EnableDebug(true)
	}
	return enableDebug
}

func (l *logger) SetOutputFile(name string) {
	dir := path.Dir(name)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		l.Error(err)
		return
	}
	l.SetWriter(file)
}

func (l *logger) SetWriter(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.impl.Out = w
}

func (l *logger) GetWriter() io.Writer {
	return l.impl.Out
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.impl.WithField("name", l.name).Info(fmt.Sprintf(format, v...))
}

func (l *logger) Debug(v ...interface{}) {
	if enableDebug {
		l.impl.WithField("name", l.name).Debug(v...)
	}
}

func (l *logger) Warn(v ...interface{}) {
	l.impl.WithField("name", l.name).Warn(v...)
}

func (l *logger) Info(v ...interface{}) {
	l.impl.WithField("name", l.name).Info(v...)
}

func (l *logger) Error(v ...interface{}) {
	l.impl.WithField("name", l.name).Error(v...)
}

func (l *logger) ErrorIf(err error, v ...interface{}) {
	if err != nil {
		a := append(append([]interface{}{}, err), v...)
		l.impl.WithField("name", l.name).Error(a...)
	}
}

func (l *logger) Fatal(v ...interface{}) {
	l.impl.WithField("name", l.name).Fatal(v...)
}

func (l *logger) Panic(v ...interface{}) {
	l.impl.WithField("name", l.name).Panic(v...)
}

type serverXContextHook struct {
}

func (hook serverXContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook serverXContextHook) Fire(entry *logrus.Entry) error {

	if enableDebug {
		if pc, file, line, ok := runtime.Caller(7); ok {
			funcName := runtime.FuncForPC(pc).Name()

			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(funcName)
			entry.Data["line"] = line
		}
	}
	return nil
}

type serverXTextFormatter struct {
}

var timeFormat = "2006-01-02 15:04:05.999999"

func (f *serverXTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	callerinfo := ""
	if enableDebug {
		callerinfo = fmt.Sprintf("\t\t%s (%s:%d)", entry.Data["func"], entry.Data["file"], entry.Data["line"])
	}

	name := entry.Data["name"]
	level := "unknown"

	switch entry.Level {
	case logrus.DebugLevel:
		level = "DEBUG"
	case logrus.InfoLevel:
		level = "INFO"
	case logrus.WarnLevel:
		level = "WARN"
	case logrus.ErrorLevel:
		level = "ERROR"
	case logrus.FatalLevel:
		level = "FATAL"
	case logrus.PanicLevel:
		level = "PANIC"
	}

	s := fmt.Sprintf("%-26s %-5s |%s| %s %s\n", entry.Time.Format(timeFormat), level, name, entry.Message, callerinfo)
	return []byte(s), nil
}

// Default Logger std

var std Logger = newStdLogger("global", os.Stdout)

func newStdLogger(name string, out io.Writer) *logger {
	l := &logger{
		impl: &logrus.Logger{
			Out:       out,
			Formatter: &serverXTextFormatter{},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
		name: name,
		mu:   &sync.Mutex{},
	}
	l.impl.AddHook(serverXContextHook1{})
	loggerManager = append(loggerManager, l)
	return l
}

type serverXContextHook1 struct {
}

func (hook serverXContextHook1) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook serverXContextHook1) Fire(entry *logrus.Entry) error {

	if enableDebug {
		if pc, file, line, ok := runtime.Caller(8); ok {
			funcName := runtime.FuncForPC(pc).Name()

			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(funcName)
			entry.Data["line"] = line
		}
	}
	return nil
}

// Printf .
func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}

// Debug .
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Warn .
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Info .
func Info(v ...interface{}) {
	std.Info(v...)
}

// Error .
func Error(v ...interface{}) {
	std.Error(v...)
}

// ErrorIf .
func ErrorIf(err error, v ...interface{}) {
	std.ErrorIf(err, v...)
}

// Fatal .
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Panic .
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// SetOutputFile .
func SetOutputFile(name string) {
	std.SetOutputFile(name)
}

// SetWriter .
func SetWriter(w io.Writer) {
	std.SetWriter(w)
}

// GetWriter .
func GetWriter(w io.Writer) {
	std.GetWriter()
}
