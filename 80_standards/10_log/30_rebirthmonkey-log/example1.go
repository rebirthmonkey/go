// this example comes initially from github.com/rebirthmonkey/go/pkg/log/examples
package main

import (
	"github.com/rebirthmonkey/go/pkg/log"
)

func main() {
	defer log.Flush()

	// Debug、Info(with field)、Warnf、Errorw使用
	log.Debug("This is a debug message")                         // the log leve should >= INFO
	log.Info("This is a info message", log.Int32("int_key", 10)) // display key:value pair
	log.Warnf("This is a formatted %s message", "warn")
}
