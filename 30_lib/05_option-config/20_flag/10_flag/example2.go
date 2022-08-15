package main

import (
	"flag"
	"fmt"
)

func main() {
	dataPath := flag.String("d", "/home/xxx/sample/", "DB data path")
	logFile := flag.String("l", "/home/xxx/sample.log", "log file")
	nowaitFlag := flag.Bool("w", false, "do not wait until operation completes")

	flag.Parse()

	var cmd string = flag.Arg(0)

	fmt.Println("cmd	  : ", cmd)
	fmt.Println("data path: ", *dataPath)
	fmt.Println("log file : ", *logFile)
	fmt.Println("nowait   : ", *nowaitFlag)

	fmt.Println("-------------------------------------------------------")

	fmt.Println("there are", flag.NArg(), "non-flag input param")
	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}
}
