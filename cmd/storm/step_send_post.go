package main

import (
	"log"

	"github.com/levigross/grequests"
)

type StepSendPost struct {
}

func NewStepSendPost() (step *StepSendPost) {
	step = new(StepSendPost)

	return step
}

func (step *StepSendPost) Run(context *Context) error {
	endpoint := context.Config.BaseUrl + context.getStringValue("URL")
	log.Printf("Step send POST to URL: %v\n", endpoint)

	resp, err := grequests.Post(endpoint,
		&grequests.RequestOptions{JSON: context.getStringValue("JSON")})

	if err != nil {
		log.Printf("Unable to make request: %v\n", err)
		context.defineValue("HTTP_STATUS_CODE", 0)
		return err
	}

	context.defineValue("HTTP_STATUS_CODE", resp.StatusCode)

	return nil
}

func (step *StepSendPost) StepId(context *Context) string {
	return "POST TO " + context.Config.BaseUrl + context.getStringValue("URL")
}
