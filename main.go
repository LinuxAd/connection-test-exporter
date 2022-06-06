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

	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0770)
}

func logDirInit() *string {
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

func testConnection(URL string) {
	// Set up loop to run test
	go func(URL string) {
		for {
			connTest(URL)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}(URL)

}

func main() {

	// Parse flags
	flag.IntVar(&interval, "interval", 2, "Time between test connections.")
	addr := flag.String("addr", ":80", "The address to expose the metrics on.")
	URL := flag.String("url", "https://bbc.co.uk", "The URL to test.")
	logBool := flag.Bool("log", false, "Whether to log to file or not.")
	file := flag.String("logfile", defaultLog, "The file to log to.")
	flag.Parse()

	logDests := []io.Writer{os.Stdout}

	if *logBool {
		// if a different path hasn't been set, compute the default path
		if *file == defaultLog {
			file = logDirInit()
		}

		// Open the log file for writing
		logFile, err := openLogFile(*file)
		if err != nil {
			log.Fatal(err)
		}

		logDests = append(logDests, logFile)
	}

	logOut := io.MultiWriter(logDests...)

	if *logBool {
		// Add line to log for each time the program is run
		infoLogger.Printf("script invoked using log file: %s", *file)
	}

	// Set up loggers for error and info
	infoLogger = log.New(logOut, "[info] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLogger = log.New(logOut, "[error] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

	infoLogger.Printf("metrics available on: %v", *addr)

	connectionInterval.Set(float64(interval))

	testConnection(*URL)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(
		http.ListenAndServe(*addr, nil),
	)

}
