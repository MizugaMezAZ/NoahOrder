package origin

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

// Entity is an logger instance
type Entity struct {
	lg *log.Logger

	lastYearDay int

	logfile *os.File

	// EnableMsg control log MsgRecv/MsgSend or not
	EnableMsg bool
	// EnableInfo control log info or not
	EnableInfo bool

	useStdout bool

	mu sync.Mutex
}

// NewEntity returns a new Entity, default EnableMsg, EnableInfo, will not outputStdout
func NewEntity() *Entity {
	e := &Entity{
		EnableMsg:  true,
		EnableInfo: true,
		useStdout:  false,
	}
	if err := e.checkNewDay(); err != nil {
		log.Fatalln(err)
	}

	return e
}

func (e *Entity) checkNewDayWithLock() {
	if e.useStdout {
		return
	}

	e.mu.Lock()
	e.checkNewDay()
	e.mu.Unlock()
}

func (e *Entity) checkNewDay() error {
	now := time.Now()
	yd := now.YearDay()
	if yd == e.lastYearDay {
		return nil
	}

	logfile, err := openLogFile(now)
	if err != nil {
		log.Println(err)
		return err
	}

	e.lastYearDay = yd

	if e.lg == nil {
		e.lg = log.New(logfile, "", log.Lshortfile|log.Ldate|log.Ltime)
	} else {
		e.lg.SetOutput(logfile)
	}

	if e.logfile != nil {
		e.logfile.Close()
	}
	e.logfile = logfile
	return nil
}

func exeName() string {
	return path.Base(strings.ReplaceAll(os.Args[0], "\\", "/"))
}

func openLogFile(now time.Time) (*os.File, error) {
	os.Mkdir("log", 0777)

	name := fmt.Sprintf("log/%v-%v.log", exeName(), time2YMD(now))

	logfile, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return logfile, err
}

const timeYMDLayout = "2006-01-02"

// time2YMD return timeYMDLayout "2006-01-02" given time.Time, for avoid recycle import, rewrite form ut package
func time2YMD(t time.Time) string {
	return t.Format(timeYMDLayout)
}
