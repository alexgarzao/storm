package main

import (
	"fmt"
	"log"
)

type StepDefineRequestTimeout struct {
	Timeout int
}

func NewStepDefineRequestTimeout(timeout int) (step *StepDefineRequestTimeout) {
	step = new(StepDefineRequestTimeout)
	step.Timeout = timeout

	return step
}

func (step *StepDefineRequestTimeout) Run(context *Context) error {
	log.Printf("Step define request timeout: %v\n", step.Timeout)

	return nil
}

func (step *StepDefineRequestTimeout) StepID(context *Context) string {
	return fmt.Sprintf("Define request timeout %v", step.Timeout)
}

func (step *StepDefineRequestTimeout) EndpointID(context *Context) string {
	return ""
}

func (step *StepDefineRequestTimeout) MustStat(context *Context) bool {
	return false
}
