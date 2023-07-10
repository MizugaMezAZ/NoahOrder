package zap

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// FourLetterLevelEncoder ...
func FourLetterLevelEncoder(l Level, enc zapcore.PrimitiveArrayEncoder) {
	var s string
	switch l {
	case DebugLevel:
		s = "DEBG"
	case InfoLevel:
		s = "INFO"
	case WarnLevel:
		s = "WARN"
	case ErrorLevel:
		s = "ERRO"
	case DPanicLevel:
		s = "DPAN"
	case PanicLevel:
		s = "PANC"
	case FatalLevel:
		s = "FATL"
	default:
		s = fmt.Sprintf("LEVEL(%d)", l)
	}

	enc.AppendString(s)
}

// LevelString ...
func LevelString(l Level) string {

	switch l {
	case DebugLevel:
		return "DEBG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERRO"
	case DPanicLevel:
		return "DPAN"
	case PanicLevel:
		return "PANC"
	case FatalLevel:
		return "FATL"
	}
	return fmt.Sprintf("LEVEL(%d)", l)
}
