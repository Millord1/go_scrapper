package logger

import (
	"log"
	"os"
)

var defaultLogFile string = "general.log"

func Log(toLog *string, file string) {
	if len(file) < 1 {
		file = defaultLogFile
	}
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// log.Fatalf("Error opening file: %v", err)
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(&toLog)
}
