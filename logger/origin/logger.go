package origin

import (
	"fmt"
	"gorder/logger"
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

var errorPrefix = color.RedString("ERRO ")

// Err output log with errorPrefix
func (e *Entity) Error(v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(errorPrefix)
	e.lg.Output(2, fmt.Sprintln(v...))
}

// Errf output log with errorPrefix in format
func (e *Entity) Errorf(format string, v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(errorPrefix)
	e.lg.Output(2, fmt.Sprintf(format, v...))
}

var infoPrefix = "INFO "

// Info output log with infoPrefix if config info enable
func (e *Entity) Info(v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(infoPrefix)
		e.lg.Output(2, fmt.Sprintln(v...))
	}
}

// Infof output log with infoPrefix in format if config info enable
func (e *Entity) Infof(format string, v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(infoPrefix)
		e.lg.Output(2, fmt.Sprintf(format, v...))
	}
}

var debugPrefix = color.BlueString("DEBG ")

// Debugf output log with debugPrefix in format if config info enable
func (e *Entity) Debug(v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(debugPrefix)
		e.lg.Output(2, fmt.Sprintln(v...))
	}
}

// Debugf output log with debugPrefix in format if config info enable
func (e *Entity) Debugf(format string, v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(debugPrefix)
		e.lg.Output(2, fmt.Sprintf(format, v...))
	}
}

var warnPrefix = color.YellowString("WARN ")

// Warn output log with warnPrefix if config info enable
func (e *Entity) Warn(v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(warnPrefix)
		e.lg.Output(2, fmt.Sprintln(v...))
	}
}

// Warnf output log with warnPrefix in format if config info enable
func (e *Entity) Warnf(format string, v ...interface{}) {
	if e.EnableInfo {
		e.checkNewDayWithLock()
		e.lg.SetPrefix(warnPrefix)
		e.lg.Output(2, fmt.Sprintf(format, v...))
	}
}

var fatalPrefix = color.RedString("FATL ")

// Fatalf output log with golderrorPrefix in format, and trig fatal os.Exit(1)
func (e *Entity) Fatalf(format string, v ...interface{}) {
	fmt.Println(v...)

	e.checkNewDayWithLock()
	e.lg.SetPrefix(fatalPrefix)
	e.lg.Output(2, fmt.Sprintf(format, v...))

	e.logfile.Write(debug.Stack())
	os.Exit(1)
}

// Fatal output log with golderrorPrefix, and trig fatal os.Exit(1)
func (e *Entity) Fatal(v ...interface{}) {
	fmt.Println(v...)

	e.checkNewDayWithLock()
	e.lg.SetPrefix(fatalPrefix)
	e.lg.Output(2, fmt.Sprintln(v...))

	e.logfile.Write(debug.Stack())
	os.Exit(1)
}

var panicPrefix = color.RedString("PANC ")

// Panicf output log with panicPrefix in format, and trig panic
func (e *Entity) Panicf(format string, v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(panicPrefix)
	s := fmt.Sprintf(format, v...)
	e.lg.Output(2, s)
	e.logfile.Write(debug.Stack())
	panic(s)
}

// Panic output log with panicPrefix in format, and trig panic
func (e *Entity) Panic(v ...interface{}) {
	e.checkNewDayWithLock()
	e.lg.SetPrefix(panicPrefix)
	s := fmt.Sprintln(v...)
	e.lg.Output(2, s)
	// e.logfile.Sync()
	e.logfile.Write(debug.Stack())
	panic(s)
}

func (e *Entity) With(v ...interface{}) logger.Logger {
	return e
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
func (e *Entity) CleanLog() {
	folder := "./log"
	e.Info("開始進行 log 删除, 日誌保留時間: ", defaultClearLogForwardTime)
	logs, err := ioutil.ReadDir(folder)
	if err != nil {
		e.Warn("读取 ./log dir fail, err:", err)
		return
	}
	for _, log := range logs {
		if time.Now().Sub(log.ModTime()) > defaultClearLogForwardTime {
			rmFile := fmt.Sprintf("%s/%s", folder, log.Name())
			e.Info("開始刪除:", rmFile, ", 最後更改時間:", log.ModTime().Format("2006/01/02 15:04:05"))
			// 文件夹
			if err := os.Remove(rmFile); err != nil {
				e.Error("刪除日誌:", rmFile, "出錯了, err:", err)
			}
		}
		// log.ModTime()
	}
	e.Info("结束 log 删除, 完工了")
}

// SetClearLogForwardTime ...
func SetClearLogForwardTime(t time.Duration) {
	defaultClearLogForwardTime = t
}

// GetClearLogForwardTime ...
func GetClearLogForwardTime() time.Duration {
	return defaultClearLogForwardTime
}
