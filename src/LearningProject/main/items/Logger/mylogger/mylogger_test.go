package mylogger

import (
	"os"
	"strings"
	"testing"
	"time"
)

// 测试 parseLogLevel 函数
func TestParseLogLevel(t *testing.T) {
	testCases := []struct {
		input    string
		expected LogLevel
	}{
		{"DEBUG", DEBUG},
		{"TRACE", TRACE},
		{"INFO", INFO},
		{"WARNING", WARNING},
		{"ERROR", ERROR},
		{"FATAL", FATAL},
		{"invalid", UNKNOWN},
	}

	for _, tc := range testCases {
		level, err := parseLogLevel(tc.input)
		if level != tc.expected {
			t.Errorf("parseLogLevel(%s) = %v; want %v", tc.input, level, tc.expected)
		}
		if tc.input != "invalid" && err != nil {
			t.Errorf("parseLogLevel(%s) returned an unexpected error: %v", tc.input, err)
		}
		if tc.input == "invalid" && err == nil {
			t.Errorf("parseLogLevel(%s) did not return an error as expected", tc.input)
		}
	}
}

// 测试 NewLog 函数和日志写入
func TestNewLogAndWrite(t *testing.T) {
	logFilePath := "test_log"
	log, err := NewLog("DEBUG", logFilePath)
	if err != nil {
		t.Fatalf("NewLog() returned an error: %v", err)
	}
	defer log.Close()

	// 写入一条日志
	log.Debug("Test log message")

	// 等待一段时间确保日志写入文件
	time.Sleep(100 * time.Millisecond)

	// 检查日志文件是否存在
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	logFileName := logFilePath + "-" + dateStr + ".log"
	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		t.Errorf("Log file %s does not exist", logFileName)
	}

	// 读取日志文件内容，检查是否包含测试日志消息
	fileContent, err := os.ReadFile(logFileName)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}
	if !strings.Contains(string(fileContent), "Test log message") {
		t.Errorf("Log file does not contain the test log message")
	}
}

// 测试日志级别过滤
func TestLogLevelFiltering(t *testing.T) {
	logFilePath := "test_log_filter"
	log, err := NewLog("INFO", logFilePath)
	if err != nil {
		t.Fatalf("NewLog() returned an error: %v", err)
	}
	defer log.Close()

	// 写入不同级别的日志
	log.Debug("Debug log message")
	log.Info("Info log message")

	// 等待一段时间确保日志写入文件
	time.Sleep(100 * time.Millisecond)

	// 读取日志文件内容，检查是否只包含 INFO 级别的日志
	now := time.Now()
	dateStr := now.Format("2006-01-02")
	logFileName := logFilePath + "-" + dateStr + ".log"
	fileContent, err := os.ReadFile(logFileName)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}
	if strings.Contains(string(fileContent), "Debug log message") {
		t.Errorf("Log file contains DEBUG log message when level is set to INFO")
	}
	if !strings.Contains(string(fileContent), "Info log message") {
		t.Errorf("Log file does not contain INFO log message")
	}
}
