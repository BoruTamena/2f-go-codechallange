package main

import (
	"log"
	"os"
)

type Logger interface {
	Info(string)
	Warning(string)
	Error(string)
}

type Log struct {
	file *os.File
}

func (lg *Log) Info(message string) {

	content := "info:" + message

	_, err := lg.file.WriteString(content)

	if err != nil {
		log.Fatal(err)

	}

}

func (lg *Log) Warning(message string) {
	content := "warning:" + message

	_, err := lg.file.WriteString(content)

	if err != nil {
		log.Fatal(err)

	}

}

func (lg *Log) Error(message string) {

	content := "error:" + message

	_, err := lg.file.WriteString(content)

	if err != nil {
		log.Fatal(err)

	}

}

func NewLog() Log {

	// creating a logger file

	file_ptr, err := os.Create("logger_file.txt")

	if err != nil {
		log.Fatal(err)
		return Log{}
	}

	log_file := Log{
		file: file_ptr,
	}

	return log_file

}

func main() {

	// creating new log file
	lg := NewLog()

	defer lg.file.Close()

	lg.Info("Information message \n")
	lg.Warning("Warning message \n")
	lg.Error("Error message \n")

}
