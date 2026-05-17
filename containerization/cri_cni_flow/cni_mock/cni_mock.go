package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logFilePath string //nolint:gochecknoglobals

func main() {
	timestamp := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("[%s] CNI CALL RECEIVED | ENVs: %v \n", timestamp, os.Environ())

	//nolint:gosec //G304
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(logEntry)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintln(os.Stdout, "{}")

	os.Exit(0)
}
