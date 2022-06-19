package main

import (
	"flag"
	"fmt"
)

func main(){
	var dataPath, logFile string
	var noWait bool

	flag.StringVar(&dataPath, "d","/home/xxx/sample/","DB data path")
	flag.StringVar(&logFile, "l","/home/xxx/sample.log","log file")
	flag.BoolVar(&noWait, "w",false,"do not wait until operation completes")

	flag.Parse()

	fmt.Println("data path: ", dataPath)
	fmt.Println("log file : ", logFile)
	fmt.Println("no wait  : ", noWait)
}

