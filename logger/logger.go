package logger

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

/*
	支援运行更改日志级别
	TODO: 支援 rotation
	TODO: 支援 logstash
	TODO: 支援多输出

*/

// DefaultLogger 整個應用程序需要一個記錄器的全局實例，因此可以在一處更改日誌配置並將其應用於整個應用程序。
var DefaultLogger Logger = &Log{}

// Logger 無需修改業務代碼即可切換到其他日誌庫 不直接依賴任何日誌庫
type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Warnf(format string, args ...interface{})
	Warn(args ...interface{})

	Debugf(format string, args ...interface{})
	Debug(args ...interface{})

	Panicf(format string, args ...interface{})
	Panic(args ...interface{})

	With(val ...interface{}) Logger
}

// SetDefaultLogger ..
func SetDefaultLogger(l Logger) {
	DefaultLogger = l
}

// GetRotationWriter ...
func GetRotationWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(filename+"_link.log"),
		rotatelogs.WithMaxAge(time.Hour*30*24),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}

	return hook
}

// Error ....
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Errorf ....
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}

// Fatalf ....
func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatalf(format, args...)
}

// Fatal ....
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Infof ....
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// Info ....
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Warnf ....
func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}

// Warn ....
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Debugf ....
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

// Debug ....
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Panicf ....
func Panicf(format string, args ...interface{}) {
	DefaultLogger.Panicf(format, args...)
}

// Panic ....
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// With ...
func With(val ...interface{}) Logger {
	return DefaultLogger.With(val...)
}
