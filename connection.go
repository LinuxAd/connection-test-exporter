package main

import (
	"fmt"
	"net/http"

	"github.com/enescakir/emoji"
)

func connTest(url string) {
	status := "fail"
	e := emoji.Parse(":rage:")
	suf := "response nil"
	log := errorLogger

	connectionAttempts.Inc()
	resp, err := http.Get(url)
	if err != nil {
		errorLogger.Println(err)
		lastConnectionSuccessful.Set(0)
	}
	if resp != nil {
		connectionSuccesses.Inc()
		lastConnectionSuccessful.Set(1)
		status = "success"
		suf = fmt.Sprintf("status code: %d", resp.StatusCode)
		e = emoji.Parse(":+1:")
		log = infoLogger
	}

	log.Printf("%v Connection to %s was a %s, %s", e, url, status, suf)
}
