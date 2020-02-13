package utils

import (
	"fmt"
	"time"
)

const withSpace = "%s "

func LogMessage(logData interface{}) {
	timeStamp := fmt.Sprintf(withSpace, time.Now().Format(time.Stamp))

	fmt.Println(fmt.Sprint(timeStamp, logData))
}

func LogMessageWithData(message string, data interface{}) {
	LogMessage(fmt.Sprintf("%s %s", message, data))
}

func ErrorLog(err error) string {
	timeStamp := fmt.Sprintf(withSpace, time.Now().Format(time.Stamp))

	return fmt.Sprintf("%s %s", timeStamp, err.Error())
}
