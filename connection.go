package main

import (
	"fmt"
	"net/http"

	"github.com/enescakir/emoji"
)

const (
	emojiBad  = ":rage:"
	emojiGood = ":+1:"
)

func connTest(url string) {
	status := "fail"
	e := emoji.Parse(emojiBad)
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
		e = emoji.Parse(emojiGood)
		log = infoLogger
	}

	log.Printf("%v Connection to %s was a %s, %s", e, url, status, suf)
}
