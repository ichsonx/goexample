package main

import (
	"log"
	"os"
)

var logger *log.Logger
var f *os.File
var err error

func main() {

	f, err = os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		logger.Fatal(err)
	}
	logger = log.New(f, "", log.LstdFlags)
	logger.SetPrefix("TEST LOG ")
	logger.SetFlags(log.LstdFlags|log.Lshortfile)

}
