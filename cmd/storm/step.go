package main

type Step interface {
	Run(*Context) error
	StepID(*Context) string
	EndpointID(*Context) string
	MustStat(*Context) bool
}
