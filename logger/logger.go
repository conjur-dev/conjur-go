package logger

import (
	"fmt"
	"time"
)

func TimestampFormat() string {
	return "2006-01-02 15:04:05.000"
}

func Debug(log interface{}) {
	fmt.Printf("[%s] [DEBUG] %v\n", time.Now().UTC().Format(TimestampFormat()), log)
}

func Info(log interface{}) {
	fmt.Printf("[%s] [INFO] %v\n", time.Now().UTC().Format(TimestampFormat()), log)
}

func Warn(log interface{}) {
	fmt.Printf("[%s] [WARN] %v\n", time.Now().UTC().Format(TimestampFormat()), log)
}

func Error(log interface{}) {
	fmt.Printf("[%s] [ERROR] %v\n", time.Now().UTC().Format(TimestampFormat()), log)
}
