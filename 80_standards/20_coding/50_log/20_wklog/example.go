



// this example initially comes from github.com/rebirthmonkey/pkg/wklog/examples/
package main

import (
	"github.com/rebirthmonkey/pkg/wklog"
	"os"
)

func main() {
	// show log in the terminal
	wklog.Debug("debug log")
	wklog.SetOptions(wklog.WithLevel(wklog.InfoLevel))
	wklog.Debug("debug log again") // cannot see since log level should >= Info
	wklog.Info("info log")
	wklog.SetOptions(wklog.WithFormatter(&wklog.JsonFormatter{IgnoreBasicFields: false}))
	wklog.Info("log in json format")
	wklog.Warn("log in json format")

	// output log to files
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		wklog.Fatal("create file test.log failed")
	}
	defer fd.Close()

	fileLog := wklog.New(wklog.WithLevel(wklog.InfoLevel),
		wklog.WithOutput(fd),
		wklog.WithFormatter(&wklog.JsonFormatter{IgnoreBasicFields: false}),
	)
	fileLog.Debug("log with debug and with json formatter")
	fileLog.Info("log with info and with json formatter")
	fileLog.Warn("log with warn and with json formatter")
}
