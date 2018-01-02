package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type StepDefineJson struct {
	Json string
}

func NewStepDefineJson(json string) (step *StepDefineJson) {
	step = new(StepDefineJson)
	step.Json = json

	return step
}

func (step *StepDefineJson) Run(context *Context) error {
	tmpl, err := template.New("test").Parse(step.Json)
	if err != nil {
		return err
	}

	var data bytes.Buffer

	if err = tmpl.Execute(&data, context); err != nil {
		return err
	}

	result := data.String()
	log.Printf("Step define JSON: %v\n", result)
	context.defineValue("JSON", result)

	return nil
}

func (step *StepDefineJson) StepID(context *Context) string {
	return fmt.Sprintf("Define JSON")
}

func (step *StepDefineJson) EndpointID(context *Context) string {
	return ""
}

func (step *StepDefineJson) MustStat(context *Context) bool {
	return false
}
