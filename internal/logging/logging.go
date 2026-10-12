package logging

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"charm.land/log/v2"
)

const (
	LevelTrace log.Level = -8
	LevelDebug           = log.DebugLevel // -4
	LevelInfo            = log.InfoLevel  // 0
	LevelWarn            = log.WarnLevel  // 4
	LevelError           = log.ErrorLevel // 8
	LevelFatal           = log.FatalLevel // 12
)

var traceLogger *log.Logger

// nolint: misspell
func init() {
	traceLogger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           LevelTrace,
		ReportTimestamp: true,
		Prefix:          lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("TRACE"),
	})
}

func ParseLevel(level string) (log.Level, error) {
	switch level {
	case "trace":
		return LevelTrace, nil
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "fatal":
		return LevelFatal, nil
	default:
		return 0, fmt.Errorf("invalid log level: %q", level)
	}
}

func SetLevel(level log.Level) {
	traceLogger.SetLevel(level)
	log.SetLevel(level)
}

func Trace(msg string) {
	traceLogger.Log(LevelTrace, msg)
}

func Tracef(format string, args ...any) {
	traceLogger.Log(LevelTrace, fmt.Sprintf(format, args...))
}

func Debug(msg string) {
	log.Log(LevelDebug, msg)
}

func Debugf(format string, args ...any) {
	log.Log(LevelDebug, fmt.Sprintf(format, args...))
}

func Info(msg string) {
	log.Log(LevelInfo, msg)
}

func Infof(format string, args ...any) {
	log.Log(LevelInfo, fmt.Sprintf(format, args...))
}

func Warn(msg string) {
	log.Log(LevelWarn, msg)
}

func Warnf(format string, args ...any) {
	log.Log(LevelWarn, fmt.Sprintf(format, args...))
}

func Error(args ...any) {
	log.Log(LevelError, fmt.Sprint(args...))
}

func Errorf(format string, args ...any) {
	log.Log(LevelError, fmt.Sprintf(format, args...))
}

func Fatal(args ...any) {
	log.Log(LevelFatal, fmt.Sprint(args...))
	os.Exit(1)
}

func Fatalf(format string, args ...any) {
	log.Log(LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}
