package mylogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// 日志级别
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger 日志结构体
type Logger struct {
	Level    LogLevel
	filePath string
	lastDate string
	file     *os.File
	logChan  chan *logMsg
}

// 日志内容结构体
type logMsg struct {
	level     LogLevel
	fileName  string
	funcName  string
	timestamp string
	lineNo    int
	msg       string
}

// parseLogLevel 字符串转日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// logLeveltoString 日志级别转字符串
func logLeveltoString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func getInfo(n int) (funcName, fileName string, lineNo int) {
	// pc表示函数指针，file表示文件名（绝对路径），lineNo表示行号，ok表示是否成功
	pc, file, lineNo, ok := runtime.Caller(n) //n表示第几个调用者，0表示自己，1表示调用自己的调用者，2表示调用者的调用者，依次类推
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //获取函数名->哪个包的哪个函数，会出现main.main
	fileName = path.Base(file)              //获取文件名
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNo
}

// NewLog 构造函数
func NewLog(levelStr string, filePath string) (*Logger, error) {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s-%s.log", filePath, dateStr)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	log := &Logger{
		Level:    level,
		filePath: filePath,
		lastDate: dateStr,
		file:     file,
		logChan:  make(chan *logMsg, 50000),
	}
	go log.writeLogBackground()
	return log, err
}

// 判断日志级别
func (l *Logger) enable(loglevel LogLevel) bool {
	return loglevel >= l.Level
}

// 关闭日志文件
func (l *Logger) Close() error {
	return l.file.Close()
}

// 检查并切换日志文件
func (l *Logger) checkAndSwitchFile() error {
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	if dateStr != l.lastDate {
		if err := l.file.Close(); err != nil {
			return err
		}
		logFilePath := fmt.Sprintf("%s-%s.log", l.filePath, dateStr)
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		l.file = file
		l.lastDate = dateStr
	}
	return nil
}
