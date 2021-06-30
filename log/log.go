package log

// levels
const (
	DebugLevel   = 0
	ReleaseLevel = 1
	ErrorLevel   = 2
	FatalLevel   = 3
)

const (
	PrintDebugLevel   = "[debug  ] "
	PrintReleaseLevel = "[release] "
	PrintErrorLevel   = "[error  ] "
	PrintFatalLevel   = "[fatal  ] "
)

type Logger interface {
	Close()
	DoPrintf(level int, printLevel string, format string, a ...interface{})
}

//type Log struct {
//	level      int
//	baseLogger *log.Logger
//	baseFile   *os.File
//}
//
//func New(strLevel string, pathname string, flag int) (Logger, error) {
//	// level
//	var level int
//	switch strings.ToLower(strLevel) {
//	case "debug":
//		level = DebugLevel
//	case "release":
//		level = ReleaseLevel
//	case "error":
//		level = ErrorLevel
//	case "fatal":
//		level = FatalLevel
//	default:
//		return nil, errors.New("unknown level: " + strLevel)
//	}
//
//	// logger
//	var baseLogger *log.Logger
//	var baseFile *os.File
//	if pathname != "" {
//		now := time.Now()
//
//		filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
//			now.Year(),
//			now.Month(),
//			now.Day(),
//			now.Hour(),
//			now.Minute(),
//			now.Second())
//
//		file, err := os.Create(path.Join(pathname, filename))
//		if err != nil {
//			return nil, err
//		}
//
//		baseLogger = log.New(file, "", flag)
//		baseFile = file
//	} else {
//		baseLogger = log.New(os.Stdout, "", flag)
//	}
//
//	// new
//	logger := new(Log)
//	logger.level = level
//	logger.baseLogger = baseLogger
//	logger.baseFile = baseFile
//
//	return logger, nil
//}
//
//// It's dangerous to call the method on logging
//func (logger *Log) Close() {
//	if logger.baseFile != nil {
//		logger.baseFile.Close()
//	}
//
//	logger.baseLogger = nil
//	logger.baseFile = nil
//}
//
//func (logger *Log) DoPrintf(level int, printLevel string, format string, a ...interface{}) {
//	if level < logger.level {
//		return
//	}
//	if logger.baseLogger == nil {
//		panic("logger closed")
//	}
//
//	format = printLevel + format
//	logger.baseLogger.Output(3, fmt.Sprintf(format, a...))
//
//	if level == FatalLevel {
//		os.Exit(1)
//	}
//}

//func (logger *Log) Debug(format string, a ...interface{}) {
//	logger.DoPrintf(DebugLevel, PrintDebugLevel, format, a...)
//}
//
//func (logger *Log) Release(format string, a ...interface{}) {
//	logger.DoPrintf(ReleaseLevel, PrintReleaseLevel, format, a...)
//}
//
//func (logger *Log) Error(format string, a ...interface{}) {
//	logger.DoPrintf(ErrorLevel, PrintErrorLevel, format, a...)
//}
//
//func (logger *Log) Fatal(format string, a ...interface{}) {
//	logger.DoPrintf(FatalLevel, PrintFatalLevel, format, a...)
//}

//var gLogger, _ = New("debug", "", log.LstdFlags)

var gLogger Logger

// It's dangerous to call the method on logging
func Export(logger Logger) {
	if logger != nil {
		gLogger = logger
	}
}

func Debug(format string, a ...interface{}) {
	gLogger.DoPrintf(DebugLevel, PrintDebugLevel, format, a...)
}

func Release(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseLevel, PrintReleaseLevel, format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.DoPrintf(ErrorLevel, PrintErrorLevel, format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.DoPrintf(FatalLevel, PrintFatalLevel, format, a...)
}

func Close() {
	gLogger.Close()
}
