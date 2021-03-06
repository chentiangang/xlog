package xlog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LogData struct {
	timeStr  string
	levelStr string
	module   string
	filename string
	funcName string
	lineNo   int
	data     string
	color    uint8
}

type XLogBase struct {
	level  int
	module string
}

func (l *XLogBase) writeLog(file *os.File, logData *LogData) {
	fmt.Fprintf(file, "\x1b[%dm%s %s %s (%s:%s:%d)\x1b[0m %s\n",
		logData.color, logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (l *XLogBase) formatLogger(level int, module string, format string, args ...interface{}) *LogData {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.000")

	levelStr, color := getLevelStr(level)
	filename, funcName, lineNo := getLineInfo(5)

	filename = filepath.Base(filename)
	data := fmt.Sprintf(format, args...)

	return &LogData{
		timeStr:  timeStr,
		levelStr: levelStr,
		module:   module,
		filename: filename,
		lineNo:   lineNo,
		funcName: funcName,
		data:     data,
		color:    color,
	}
}

func getLineInfo(skip int) (filename, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		fun := runtime.FuncForPC(pc)
		funcName = fun.Name()
	}
	filename = file
	lineNo = line
	return
}
