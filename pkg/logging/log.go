package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	//0 父级 1父级的父级 2父级的父级的父级 依次类推
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	filePath := getLogFilePath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// 设置前缀
func setPrefix(level Level) {
	//DefaultCallerDepth 0 报告Caller()的调用者的信息
	//					 1 报告Caller()的调用者的调用者的信息
	//					 2 报告Caller()的调用者的调用者的调用者的信息
	//返回值
	//_ 调用栈标识符
	//file 带路径的完整文件名
	//line 调用在文件中的行号
	//ok 是否可以获得信息
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}
