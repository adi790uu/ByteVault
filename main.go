package main

import (
	"fmt"
	"os"
)

func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
}

func LogAppend(fp *os.File, line string) error {
	buf := []byte(line)
	buf = append(buf, '\n')

	_, err := fp.Write(buf)
	if err != nil {
		return err
	}

	return fp.Sync()
}

func main() {
 
    logPath := "log.txt"

    file, err := LogCreate(logPath)
    if err != nil {
        fmt.Printf("Error creating log file: %v\n", err)
        return
    }
    defer file.Close() 

    if err := LogAppend(file, "First log entry"); err != nil {
        fmt.Printf("Error appending to log file: %v\n", err)
        return
    }
    if err := LogAppend(file, "Second log entry"); err != nil {
        fmt.Printf("Error appending to log file: %v\n", err)
        return
    }

    fmt.Println("Log entries appended successfully.")
}