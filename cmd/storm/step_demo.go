package main

import "log"

type StepDemo struct {
	Type int //TODO
}

func NewStepDemo() (step *StepDemo) {
	step = new(StepDemo)
	// scenary.Steps = 0 // Precisa?

	return step
}

func (step *StepDemo) Run() {
	if step.Type == 1 {
		log.Printf("Step demo type 1\n")
	}

	// return nil
}
