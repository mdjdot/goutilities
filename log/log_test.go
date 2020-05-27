package log

import (
	"syscall"
	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog("./app.log")
	defer Close()

	Logger.Println("log a log")
	signalChan <- syscall.SIGTERM
}
