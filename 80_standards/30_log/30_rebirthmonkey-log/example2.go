// this example comes initially from github.com/rebirthmonkey/pkg/log/examples
package main

import (
	"context"
	"flag"

	"github.com/rebirthmonkey/pkg/log"
)

var (
	helpFlag   bool
	levelFlag  string
	formatFlag string
)

func main() {
	flag.BoolVar(&helpFlag, "h", false, "Print this help.")
	flag.StringVar(&levelFlag, "l", "debug", "Log level.")
	flag.StringVar(&formatFlag, "f", "console", "log output format.")

	flag.Parse()

	if helpFlag {
		flag.Usage()
		return
	}

	// logger configuration
	opts := &log.Options{
		Level:            levelFlag,
		Format:           formatFlag,
		EnableColor:      true, // if you need output to local path, with EnableColor must be false.
		DisableCaller:    true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
	}
	// 初始化全局logger
	log.Init(opts)
	defer log.Flush()

	// Debug、Info(with field)、Warnf、Errorw使用
	log.Debug("Debug message")
	log.Info("Info message", log.Int32("int_key", 10)) // show key/value pair
	log.Warnf("Warn message formatted with %s message", "warn")
	// Errorw() automatically combine key:value pairs
	log.Errorw("Error message printed with Errorw() function", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")

	// WithValues: automatically attach key/value pair in the end of each log
	lv := log.WithValues("X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	lv.Infow("Info message printed with [WithValues] logger")
	lv.Debugw("Debug message printed with [WithValues] logger")

	// Context: a set of key/value pairs
	ctx := lv.WithContext(context.Background())
	lc := log.FromContext(ctx)
	lc.Info("Info message printed with [WithContext] logger")

	ln := lv.WithName("NAME")
	ln.Info("Info message printed with [WithName] logger")

	// V level: pre-fix log level, then use Info() at all
	log.V(log.InfoLevel).Info("This is a InfoLevel message with Info() function")
	log.V(log.WarnLevel).Info("This is a WarnLevel message with Info() function")
	log.V(log.ErrorLevel).Info("This is a ErrorLevel message with Info() function")
}
