package main

import (
	"log"
	"os"
)

func main(){
	fileName := "infoFirst.log"
	logFile,err  := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}

	debugLog := log.New(logFile, "[Info]", log.Ltime)
	debugLog.Println("A Info message here")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug Message here ")
}
