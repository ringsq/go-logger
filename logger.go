package log

// RootLogger is the logger used for top-level messages
var RootLogger Logger

func init() {
	RootLogger = NewSimple()
}

// Debug log message
func Debug(msg ...interface{}) {
	RootLogger.Debug(msg)
}

// Info log message
func Info(msg ...interface{}) {
	RootLogger.Info(msg)
}

// Warn log message
func Warn(msg ...interface{}) {
	RootLogger.Warn(msg)
}

// Error log message
func Error(msg ...interface{}) {
	RootLogger.Error(msg)
}

// Fatal log message (and exit)
func Fatal(msg ...interface{}) {
	RootLogger.Fatal(msg)
}

// Panic log message (and exit)
func Panic(msg ...interface{}) {
	RootLogger.Panic(msg)
}

// Debugln message with linefeed
func Debugln(msg ...interface{}) {
	RootLogger.Debugln(msg)
}

// Infoln message with linefeed
func Infoln(msg ...interface{}) {
	RootLogger.Infoln(msg)
}

// Warnln message with linefeed
func Warnln(msg ...interface{}) {
	RootLogger.Warnln(msg)
}

// Errorln message with linefeed
func Errorln(msg ...interface{}) {
	RootLogger.Errorln(msg)
}

// Fatalln log message (and exit)
func Fatalln(msg ...interface{}) {
	RootLogger.Fatalln(msg)
}

// Panicln log message (and exit)
func Panicln(msg ...interface{}) {
	RootLogger.Panicln(msg)
}

// Debugf formatted message
func Debugf(format string, args ...interface{}) {
	RootLogger.Debugf(format, args)
}

// Infof formatted message
func Infof(format string, args ...interface{}) {
	RootLogger.Infof(format, args)
}

// Warnf formatted message
func Warnf(format string, args ...interface{}) {
	RootLogger.Warnf(format, args)
}

// Errorf formatted message
func Errorf(format string, args ...interface{}) {
	RootLogger.Errorf(format, args)
}

// Fatalf log message (and exit)
func Fatalf(format string, args ...interface{}) {
	RootLogger.Fatalf(format, args)
}

// Panicf formatted panic message
func Panicf(format string, args ...interface{}) {
	RootLogger.Fatalf(format, args)
}

// WithFields will return a new logger based on the original logger
// with the additional supplied fields
func WithFields(fields Fields) Logger {
	return RootLogger.WithFields(fields)
}
