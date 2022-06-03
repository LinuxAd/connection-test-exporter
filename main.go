package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/enescakir/emoji"
)

const (
	defaultLog = "<HOME_DIR>/connection-tests/<date>.log"
	defaultDir = "connection-tests"
)

var (
	infoLogger  = &log.Logger{}
	errorLogger = &log.Logger{}
	interval    = 0
)

func openLogFile(path string) (io.Writer, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	mw := io.MultiWriter(os.Stdout, logFile)

	return mw, err
}

func connTest(url string) {
	status := "fail"
	e := emoji.Parse(":rage:")
	suf := "response nil"
	log := errorLogger

	resp, err := http.Get(url)
	if err != nil {
		errorLogger.Println(err)
	}
	if resp != nil {
		status = "success"
		suf = fmt.Sprintf("status code: %d", resp.StatusCode)
		e = emoji.Parse(":+1:")
		log = infoLogger
	}

	log.Printf("%v Connection to %s was a %s, %s", e, url, status, suf)
}

func logInit() *string {
	// Find home directory
	x, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find your home directory")
	}

	// Get neat folder to store logs
	y := filepath.Join(x, defaultDir)
	// Make neat folder
	if err := os.Mkdir(y, 0644); err != nil {
		if !os.IsExist(err) {
			// If the error isn't that the directory already exists, log the fatal error. Otherwise move on
			log.Fatalf("could not make the %s directory: %s", defaultDir, err.Error())
		}
	}

	// Add new file name per day
	t := time.Now().Format("Mon-2-Jan-2006")
	fullPath := filepath.Join(y, t)

	// Append the ".log" file extension to path
	final := fmt.Sprintf("%s.log", fullPath)
	return &final
}

func main() {

	// Parse flags
	flag.IntVar(&interval, "interval", 2, "Time between test connections.")
	URL := flag.String("url", "https://bbc.co.uk", "The URL to test.")
	file := flag.String("log", defaultLog, "The file to log to.")
	flag.Parse()

	// if a different path hasn't been set, compute the default path
	if *file == defaultLog {
		file = logInit()
	}

	// Open the log file for writing
	logFile, err := openLogFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	// Set up loggers for error and info
	infoLogger = log.New(logFile, "[info] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLogger = log.New(logFile, "[error] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

	// Add line to log for each time the program is run
	infoLogger.Printf("script invoked using log file: %s", *file)

	// Set up loop to run test
	for {
		connTest(*URL)
		time.Sleep(time.Duration(interval) * time.Second)
	}

}
