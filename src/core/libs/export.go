package libs

import (
	"core/libs/common"
	"core/libs/logger"
)

func init() {
}

//------------------------------------------------ 严重程度由高到低
func ERR(v ...interface{}) {
	logger.Error(v)
}

func WARN(v ...interface{}) {
	logger.Warn(v)
}

func INFO(v ...interface{}) {
	logger.Info(v)
}

func NOTICE(v ...interface{}) {
	logger.Notice(v)
}

func DEBUG(v ...interface{}) {
	logger.Debug(v)
}

func Run() {
	common.Run()
}

func NumToString(num interface{}) string {
	return common.NumToString(num)
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	return common.If(condition, trueVal, falseVal)
}
