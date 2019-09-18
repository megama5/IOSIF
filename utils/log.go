package utils

import (
	"fmt"
	"time"
)

const withSpace = "%s "

func Log(logData interface{}) {
	timeStamp := fmt.Sprintf(withSpace, time.Now().Format(time.Stamp))

	fmt.Println(fmt.Sprint(timeStamp, logData))
}
