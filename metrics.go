package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	metricsNamespace = "connection"
	metricsSubsys    = "tests"
)

var (
	connectionSuccesses = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsys,
		Name:      "succeeded",
		Help:      "The total number of successful connection attempts.",
	})
	lastConnectionSuccessful = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsys,
		Name:      "last_successful",
		Help:      "Records 1 or 0 depending on whether the last connection attempt was a success.",
	})
	connectionAttempts = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsys,
		Name:      "made",
		Help:      "The total number of connection attempts made.",
	})
	connectionInterval = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsys,
		Name:      "interval",
		Help:      "The value of the connection interval setting.",
	})
)
