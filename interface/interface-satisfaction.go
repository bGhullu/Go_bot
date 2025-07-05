package main

import "fmt"

type Logger interface {
	Log()
}

type MyLogger struct{}

func (l *MyLogger) Log() {
	fmt.Println("Logging via pointer receiver!")
}

func main() {
	var l Logger

	ml := MyLogger{}

	l = &ml
	l.Log()
}
