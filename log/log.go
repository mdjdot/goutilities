package log

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	// Logger 导出的 Logger
	Logger     *log.Logger
	signalChan = make(chan os.Signal) // signalChan 系统信号
	exitChan   = make(chan bool)      // exitChan 退出管道
)

// InitLog 配置日志
func InitLog(logPath string) {
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	Logger = log.New(f, "", log.LstdFlags|log.Lshortfile)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		Logger.Println("Closed by:", <-signalChan)
		f.Close()
		exitChan <- true
	}()
}

// Close 关闭日志文件
func Close() {
	<-exitChan
}
