package mylogger

// import (
// 	"fmt"
// 	"time"
// )

// func (l *Logger) writeLogBackground() {
// 	for {
// 		if err := l.checkAndSwitchFile(); err != nil {
// 			fmt.Printf("Failed to switch log file: %v\n", err)
// 			return
// 		}
// 		select {
// 		case logInfo := <-l.logChan:
// 			level := logLeveltoString(logInfo.level)
// 			logMsg := fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", logInfo.timestamp, level, logInfo.fileName, logInfo.funcName, logInfo.lineNo, logInfo.msg)
// 			_, err := l.file.WriteString(logMsg)
// 			if err != nil {
// 				fmt.Printf("Failed to write log: %v\n", err)
// 			}
// 		default:
// 			//取不到日志先休息500毫秒
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}
// }

import (
	"bufio"
	"fmt"
	"time"
)

// 使用缓冲区写入文件
func (l *Logger) writeLogBackground() {
	writer := bufio.NewWriter(l.file)
	defer writer.Flush()

	for {
		if err := l.checkAndSwitchFile(); err != nil {
			fmt.Printf("Failed to switch log file: %v\n", err)
			return
		}
		select {
		case logInfo := <-l.logChan:
			level := logLeveltoString(logInfo.level)
			logMsg := fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", logInfo.timestamp, level, logInfo.fileName, logInfo.funcName, logInfo.lineNo, logInfo.msg)
			_, err := writer.WriteString(logMsg)
			if err != nil {
				fmt.Printf("Failed to write log: %v\n", err)
			}
			// 定期刷新缓冲区
			if writer.Buffered() >= 4096 {
				writer.Flush()
			}
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 记录日志内容
func (l *Logger) recordlog(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	//获取记录的行号
	funcName, fileName, lineNo := getInfo(3)
	//1、logMsg对象
	logInfo := logMsg{
		level:     lv,
		fileName:  fileName,
		funcName:  funcName,
		timestamp: now.Format("2006-01-02 15:04:05"),
		lineNo:    lineNo,
		msg:       msg,
	}
	//2、将logMsg对象写入channel通道中
	select {
	case l.logChan <- &logInfo:
	default:
		//把日志信息丢了，保证不出现堵塞
	}
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		l.recordlog(DEBUG, format, a...)
	}
}

func (l *Logger) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		l.recordlog(TRACE, format, a...)
	}
}

func (l *Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		l.recordlog(INFO, format, a...)
	}
}

func (l *Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		l.recordlog(WARNING, format, a...)
	}
}

func (l *Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		l.recordlog(ERROR, format, a...)
	}
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		l.recordlog(FATAL, format, a...)
	}
}
