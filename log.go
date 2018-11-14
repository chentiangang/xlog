package xlog

var logger XLog = newXLog("console", XLogLevelDebug, "", "default")

type XLog interface {
	Init() error
	LogDebug(fmt string, args ...interface{})
	LogTrace(fmt string, args ...interface{})
	LogInfo(fmt string, args ...interface{})
	LogWarn(fmt string, args ...interface{})
	LogError(fmt string, args ...interface{})
	LogFatal(fmt string, args ...interface{})

	Close()
	SetLevel(level int)
}

func newXLog(logType string, level int, filename, module string) XLog {
	var logger XLog
	switch logType {
	case "file","File":
		logger = NewXFile(level, filename, module)
	case "console","Console":
		logger = NewXConsole(level, module)
	default:
		logger = NewXFile(level, filename, module)
	}
	return logger
}

func Init(logType string, level int, filename, module string) error {
	logger = newXLog(logType, level, filename, module)
	return logger.Init()
}

func LogDebug(format string, args ...interface{}) {
	logger.LogDebug(format, args...)
}

func LogTrace(format string, args ...interface{}) {
	logger.LogTrace(format, args...)
}

func LogInfo(format string, args ...interface{}) {
	logger.LogInfo(format, args...)
}

func LogWarn(format string, args ...interface{}) {
	logger.LogWarn(format, args...)
}

func LogError(format string, args ...interface{}) {
	logger.LogError(format, args...)
}

func LogFatal(format string, args ...interface{}) {
	logger.LogFatal(format, args...)
}

func Close() {
	logger.Close()
}

func SetLevel(level int) {
	logger.SetLevel(level)
}
