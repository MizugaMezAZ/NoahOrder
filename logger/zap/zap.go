package zap

import (
	"gorder/logger"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level ...
type Level = zapcore.Level

const (
	// DebugLevel  	-1
	DebugLevel Level = zap.DebugLevel
	// InfoLevel 	0
	InfoLevel Level = zap.InfoLevel
	// WarnLevel 	1
	WarnLevel Level = zap.WarnLevel
	// ErrorLevel 	2
	ErrorLevel Level = zap.ErrorLevel
	// DPanicLevel 	3
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel   4
	PanicLevel = zapcore.PanicLevel
	// FatalLevel 	5
	FatalLevel = zapcore.FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

// Field ...
type Field = zap.Field

var (
	// 这边引用 所有zap的filed方法

	// Skip       .
	Skip = zap.Skip
	// Binary     .
	Binary = zap.Binary
	// Bool       .
	Bool = zap.Bool
	// Boolp      .
	Boolp = zap.Boolp
	// ByteString .
	ByteString = zap.ByteString
	// Complex128 .
	Complex128 = zap.Complex128
	// Complex128p .
	Complex128p = zap.Complex128p
	// Complex64  .
	Complex64 = zap.Complex64
	// Complex64p .
	Complex64p = zap.Complex64p
	// Float64    .
	Float64 = zap.Float64
	// Float64p   .
	Float64p = zap.Float64p
	// Float32    .
	Float32 = zap.Float32
	// Float32p   .
	Float32p = zap.Float32p
	// Int        .
	Int = zap.Int
	// Intp       .
	Intp = zap.Intp
	// Int64      .
	Int64 = zap.Int64
	// Int64p     .
	Int64p = zap.Int64p
	// Int32      .
	Int32 = zap.Int32
	// Int32p     .
	Int32p = zap.Int32p
	// Int16      .
	Int16 = zap.Int16
	// Int16p     .
	Int16p = zap.Int16p
	// Int8       .
	Int8 = zap.Int8
	// Int8p      .
	Int8p = zap.Int8p
	// String     .
	String = zap.String
	// Stringp    .
	Stringp = zap.Stringp
	// Uint       .
	Uint = zap.Uint
	// Uintp      .
	Uintp = zap.Uintp
	// Uint64     .
	Uint64 = zap.Uint64
	// Uint64p    .
	Uint64p = zap.Uint64p
	// Uint32     .
	Uint32 = zap.Uint32
	// Uint32p    .
	Uint32p = zap.Uint32p
	// Uint16     .
	Uint16 = zap.Uint16
	// Uint16p    .
	Uint16p = zap.Uint16p
	// Uint8      .
	Uint8 = zap.Uint8
	// Uint8p     .
	Uint8p = zap.Uint8p
	// Uintptr    .
	Uintptr = zap.Uintptr
	// Uintptrp   .
	Uintptrp = zap.Uintptrp
	// Reflect    .
	Reflect = zap.Reflect
	// Namespace  .
	Namespace = zap.Namespace
	// Stringer   .
	Stringer = zap.Stringer
	// Time       .
	Time = zap.Time
	// Timep      .
	Timep = zap.Timep
	// Stack      .
	Stack = zap.Stack
	// StackSkip  .
	StackSkip = zap.StackSkip
	// Duration   .
	Duration = zap.Duration
	// Durationp  .
	Durationp = zap.Durationp
	// Object     .
	Object = zap.Object
	// Inline     .
	Inline = zap.Inline
	// Any        .
	Any = zap.Any
)

// Option .
type Option = zap.Option

// WithCaller .
var WithCaller = zap.WithCaller

// AddCallerSkip ..
var AddCallerSkip = zap.AddCallerSkip

// WithFatalHook ...
var WithFatalHook = zap.WithFatalHook

// Hooks ...
var Hooks = zap.Hooks

// Logger ...
type Logger struct {
	l        *zap.SugaredLogger
	level    Level
	withSkip bool
}

var hookWG sync.WaitGroup

// New ...
func New(writer io.Writer, level Level, debugFormat bool) *Logger {
	if writer == nil {
		panic("new logger, nil writer")
	}

	if level < _minLevel || level > _maxLevel {
		panic("out of range level")
	}

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	encoder := zapcore.NewJSONEncoder(cfg.EncoderConfig)

	if debugFormat {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = FourLetterLevelEncoder
		cfg.EncoderConfig.ConsoleSeparator = " | "
		cfg.EncoderConfig.FunctionKey = "func"
		encoder = zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)

	logger := &Logger{
		l:     zap.New(core, WithCaller(true), AddCallerSkip(2)).Sugar(),
		level: level,
	}

	logger.WithOption(WithFatalHook(logger))

	return logger
}

// Error ...
func (l *Logger) Error(args ...interface{}) {
	l.l.Error(args...)
}

// Errorf ...
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.l.Errorf(format, args...)
}

// Fatalf ...
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.l.Fatalf(format, args...)
}

// Fatal ...
func (l *Logger) Fatal(args ...interface{}) {
	l.l.Fatal(args...)
}

// Infof ...
func (l *Logger) Infof(format string, args ...interface{}) {
	l.l.Infof(format, args...)
}

// Info ...
func (l *Logger) Info(args ...interface{}) {
	l.l.Info(args...)
}

// Warnf ...
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.l.Warnf(format, args...)
}

// Warn ...
func (l *Logger) Warn(args ...interface{}) {
	l.l.Warn(args...)
}

// Debugf ...
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.l.Debugf(format, args...)
}

// Debug ...
func (l *Logger) Debug(args ...interface{}) {
	l.l.Debug(args...)
}

// Panicf ...
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.l.Panicf(format, args...)
}

// Panic ...
func (l *Logger) Panic(args ...interface{}) {
	l.l.Panic(args...)
}

// WithOption ...
func (l *Logger) WithOption(opt ...Option) *Logger {
	l.l = l.l.WithOptions(opt...)
	return l
}

// OnWrite .
func (l *Logger) OnWrite(e *zapcore.CheckedEntry, f []Field) {
	hookWG.Wait()
	os.Exit(1)
}

// With ...
func (l *Logger) With(args ...interface{}) logger.Logger {
	if l.withSkip {
		return &Logger{
			l:        l.l.With(args...),
			level:    l.level,
			withSkip: true,
		}
	}

	return &Logger{
		l:        l.l.With(args...).WithOptions(AddCallerSkip(-1)),
		level:    l.level,
		withSkip: true,
	}
}
