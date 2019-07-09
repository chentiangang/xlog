package xlog

var logger XLog = newXLog("console", XLogLevelDebug, "", "default",false)

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

func newXLog(logType string, level int, filename, module string,split bool) XLog {
	var logger XLog
	switch logType {
	case "file","File":
		logger = NewXFile(level, filename, module, split)
	case "console","Console":
		logger = NewXConsole(level, module)
	default:
		logger = NewXFile(level, filename, module,split)
	}
	return logger
}

func Init(logType string, level string, filename, module string,split bool) error {
	var l int
	switch level {
	case "debug":
		l = XLogLevelDebug
	case "trace":
		l = XLogLevelTrace
	case "info":
		l = XLogLevelInfo
	case "warn":
		l = XLogLevelWarn
	case "error":
		l = XLogLevelError
	case "fatal":
		l = XLogLevelFatal
	default:
		l = XLogLevelDebug
	}

	logger = newXLog(logType, l , filename, module,split)
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
