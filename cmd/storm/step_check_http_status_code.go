package main

import (
	"errors"
	"fmt"
	"log"
)

type StepCheckHttpStatusCode struct {
	HttpStatusCode int
}

func NewStepCheckHttpStatusCode(httpStatusCode int) (step *StepCheckHttpStatusCode) {
	step = new(StepCheckHttpStatusCode)
	step.HttpStatusCode = httpStatusCode

	return step
}

func (step *StepCheckHttpStatusCode) Run(context *Context) error {
	log.Printf("Step check http status code is %v\n", step.HttpStatusCode)

	if step.HttpStatusCode != context.getValue("HTTP_STATUS_CODE") {
		log.Printf("Status NOK\n")
		return errors.New("Status NOK")
	}

	log.Printf("Status OK\n")

	return nil
}

func (step *StepCheckHttpStatusCode) StepID(context *Context) string {
	return fmt.Sprintf("Check http status code is %v", step.HttpStatusCode)
}

func (step *StepCheckHttpStatusCode) EndpointID(context *Context) string {
	return ""
}

func (step *StepCheckHttpStatusCode) MustStat(context *Context) bool {
	return false
}
