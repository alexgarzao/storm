package main

import "time"

type Stats struct {
	Id       string
	Duration time.Duration
	Status   bool
	// AmountOfBytes int
	// NumberOfExecutions int
	// NumberOferrors     int
}
