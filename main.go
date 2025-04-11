package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type sizeWriter struct {
	w io.Writer
	n int
}

func (sw *sizeWriter) Write(p []byte) (int, error) {
	n, err := sw.w.Write(p)
	sw.n += n
	return n, err
}

func (sw *sizeWriter) Size() int {
	return sw.n
}

func getInputSize() int {
	mbInput := os.Args[1]
	mbSize, err := strconv.Atoi(mbInput)
	if err != nil || mbSize <= 0 {
		fmt.Println("Please provide a valid positive integer for size in MB.")
		os.Exit(1)
	}
	return mbSize * 1024 * 1024
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <size_in_MB>")
		return
	}

	size := getInputSize()
	w := &sizeWriter{w: os.Stdout}
	logMessage := strings.Repeat("A", 100)
	handler := slog.NewTextHandler(w, nil)
	logger := slog.New(handler)

	for w.Size() < size {
		logger.Info(logMessage)
	}
}
