package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
)

var e = NewEntity()

// DisableMsg disable entity msg log
func DisableMsg() {
	e.EnableMsg = false
}

// EnableMsg enable entity msg log
func EnableMsg() {
	e.EnableMsg = true
}

// DisableInfo disable entity info log
func DisableInfo() {
	e.EnableInfo = false
}

// EnableInfo enable entity info log
func EnableInfo() {
	e.EnableInfo = true
}

// var recvPrefix = color.HiGreenString("recv ")

// //MsgRecv output log with recvPrefix if config msg enable
// func MsgRecv(v ...interface{}) {
// 	if e.EnableMsg {
// 		e.checkNewDayWithLock()
// 		e.lg.SetPrefix(recvPrefix)
// 		e.lg.Output(2, fmt.Sprintln(v...))
// 	}
// }

// // MsgRecvf output log with recvPrefix in format if config msg enable
// func MsgRecvf(format string, v ...interface{}) {
// 	if e.EnableMsg {
// 		e.checkNewDayWithLock()
// 		e.lg.SetPrefix(recvPrefix)
// 		e.lg.Output(2, fmt.Sprintf(format, v...))
// 	}
// }

// var sendPrefix = color.HiYellowString("send ")

// // MsgSend output log with sendPrefix if config msg enable
// func MsgSend(v ...interface{}) {
// 	if e.EnableMsg {
// 		e.checkNewDayWithLock()
// 		e.lg.SetPrefix(sendPrefix)
// 		e.lg.Output(2, fmt.Sprintln(v...))
// 	}
// }

// // MsgSendf output log with sendPrefix in format if config msg enable
// func MsgSendf(format string, v ...interface{}) {
// 	if e.EnableMsg {
// 		e.checkNewDayWithLock()
// 		e.lg.SetPrefix(sendPrefix)
// 		e.lg.Output(2, fmt.Sprintf(format, v...))
// 	}
// }

var errorPrefix = color.RedString("ERRO ")

// Err output log with errorPrefix
func Error(v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(errorPrefix)
	e.lg.Output(2, fmt.Sprintln(v...))
}

// Errf output log with errorPrefix in format
func Errorf(format string, v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(errorPrefix)
	e.lg.Output(2, fmt.Sprintf(format, v...))
}

var infoPrefix = "INFO "

// Info output log with infoPrefix if config info enable
func Info(v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(infoPrefix)
		e.lg.Output(2, fmt.Sprintln(v...))
	}
}

// Infof output log with infoPrefix in format if config info enable
func Infof(format string, v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(infoPrefix)
		e.lg.Output(2, fmt.Sprintf(format, v...))
	}
}

var debugPrefix = color.BlueString("DEBG ")

// Debugf output log with debugPrefix in format if config info enable
func Debug(v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(debugPrefix)
		e.lg.Output(2, fmt.Sprintln(v...))
	}
}

// Debugf output log with debugPrefix in format if config info enable
func Debugf(format string, v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(debugPrefix)
		e.lg.Output(2, fmt.Sprintf(format, v...))
	}
}

var warnPrefix = color.YellowString("WARN ")

// Warn output log with warnPrefix if config info enable
func Warn(v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(warnPrefix)
		e.lg.Output(2, fmt.Sprintln(v...))
	}
}

// Warnf output log with warnPrefix in format if config info enable
func Warnf(format string, v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(warnPrefix)
		e.lg.Output(2, fmt.Sprintf(format, v...))
	}
}

// var golderrorPrefix = color.RedString("golderror ")

// // GoldErr output log with golderrorPrefix
// func GoldError(v ...interface{}) {
// 	e.checkNewDayWithLock()
// 	e.lg.SetPrefix(golderrorPrefix)
// 	msg := fmt.Sprintln(v...)
// 	e.lg.Output(2, msg)
// }

// // GoldErrf output log with golderrorPrefix in format
// func GoldErrorf(format string, v ...interface{}) {
// 	e.checkNewDayWithLock()
// 	e.lg.SetPrefix(golderrorPrefix)
// 	msg := fmt.Sprintf(format, v...)
// 	e.lg.Output(2, msg)
// }

var fatalPrefix = color.RedString("FATL ")

// Fatalf output log with golderrorPrefix in format, and trig fatal os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	fmt.Println(v...)

	e.checkNewDayWithLock()
	e.lg.SetPrefix(fatalPrefix)
	e.lg.Output(2, fmt.Sprintf(format, v...))

	e.logfile.Write(debug.Stack())
	os.Exit(1)
}

// Fatal output log with golderrorPrefix, and trig fatal os.Exit(1)
func Fatal(v ...interface{}) {
	fmt.Println(v...)

	e.checkNewDayWithLock()
	e.lg.SetPrefix(fatalPrefix)
	e.lg.Output(2, fmt.Sprintln(v...))

	e.logfile.Write(debug.Stack())
	os.Exit(1)
}

var panicPrefix = color.RedString("PANC ")

// Panicf output log with panicPrefix in format, and trig panic
func Panicf(format string, v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(panicPrefix)
	s := fmt.Sprintf(format, v...)
	e.lg.Output(2, s)
	e.logfile.Write(debug.Stack())
	panic(s)
}

// Panic output log with panicPrefix in format, and trig panic
func Panic(v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(panicPrefix)
	s := fmt.Sprintln(v...)
	e.lg.Output(2, s)
	// e.logfile.Sync()
	e.logfile.Write(debug.Stack())
	panic(s)
}

// PrintStack for debug trace code stack
func PrintStack() {
	e.logfile.Write(debug.Stack())
}

// SetStdoutput set entity Stdout enable
func SetStdoutput() {
	e.useStdout = true
	e.lg.SetOutput(os.Stdout)
}

var defaultClearLogForwardTime time.Duration = 60 * 24 * time.Hour

// CleanLog ...
func CleanLog() {
	folder := "./log"
	Info("開始進行 log 删除, 日誌保留時間: ", defaultClearLogForwardTime)
	logs, err := ioutil.ReadDir(folder)
	if err != nil {
		Warn("读取 ./log dir fail, err:", err)
		return
	}
	for _, log := range logs {
		if time.Now().Sub(log.ModTime()) > defaultClearLogForwardTime {
			rmFile := fmt.Sprintf("%s/%s", folder, log.Name())
			Info("開始刪除:", rmFile, ", 最後更改時間:", log.ModTime().Format("2006/01/02 15:04:05"))
			// 文件夹
			if err := os.Remove(rmFile); err != nil {
				Error("刪除日誌:", rmFile, "出錯了, err:", err)
			}
		}
		// log.ModTime()
	}
	Info("结束 log 删除, 完工了")
}

// SetClearLogForwardTime ...
func SetClearLogForwardTime(t time.Duration) {
	defaultClearLogForwardTime = t
}

// GetClearLogForwardTime ...
func GetClearLogForwardTime() time.Duration {
	return defaultClearLogForwardTime
}
