package main

type Step interface {
	Run(*Context) error
	StepId(*Context) string
	// NeedStat() bool indica se deve mensurar para estatisticas (de tempo? e se der erro?)
}
