package main

import "time"

// const (
// 	Scenario = iota
// 	Endpoint = iota
// )

type Stats struct {
	ScenarioID string
	EndpointID string
	Duration   time.Duration
	Status     bool
	MustStat   bool
	// AmountOfBytes int
	// NumberOfExecutions int
	// NumberOferrors     int
}
