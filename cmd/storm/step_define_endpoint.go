package main

import (
	"fmt"
	"log"
)

type StepDefineEndpoint struct {
	Url string
}

func NewStepDefineEndpoint(url string) (step *StepDefineEndpoint) {
	step = new(StepDefineEndpoint)
	step.Url = url

	return step
}

func (step *StepDefineEndpoint) Run(context *Context) error {
	log.Printf("Step define URL: %v\n", step.Url)
	context.defineValue("URL", step.Url)

	return nil
}

func (step *StepDefineEndpoint) StepID(context *Context) string {
	return fmt.Sprintf("Define URL %v", step.Url)
}

func (step *StepDefineEndpoint) EndpointID(context *Context) string {
	return ""
}

func (step *StepDefineEndpoint) MustStat(context *Context) bool {
	return false
}
