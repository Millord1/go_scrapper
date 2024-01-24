package logger_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLogFile string = "logger_test.log"

func TestLog(t *testing.T) {
	f, err := os.OpenFile(testLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	assert.Nil(t, err)
	defer f.Close()

	log.SetOutput(f)
	log.Println("Test logging")

	fi, err := os.Stat(testLogFile)

	assert.NotEqual(t, fi.Size(), 0)
}
