package main

import (
	"fmt"
	"net/http"

	"github.com/enescakir/emoji"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	connectionSuccesses = promauto.NewCounter(prometheus.CounterOpts{
		Name: "connection_tests_succeeded",
		Help: "The total number of successful connection attempts.",
	})
)

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
		connectionSuccesses.Inc()
		status = "success"
		suf = fmt.Sprintf("status code: %d", resp.StatusCode)
		e = emoji.Parse(":+1:")
		log = infoLogger
	}

	log.Printf("%v Connection to %s was a %s, %s", e, url, status, suf)
}
