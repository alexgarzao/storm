package main

import (
	"fmt"
	"log"
	"time"
)

type LoadTest struct {
	Steps    int
	Config   *Config
	UniqueId int
}

func NewLoadTest(config *Config) (loadTest *LoadTest) {
	loadTest = new(LoadTest)
	loadTest.Config = config
	loadTest.Steps = 0 // Precisa?

	return loadTest
}

func (loadTest *LoadTest) Run(scenary *Scenary) error {
	stats := make(chan Stats, loadTest.Config.UniqueIds)

	loadTest.Config.wg.Add(loadTest.Config.UniqueIds)

	statsDone := make(chan bool, 1)

	go collectStats(loadTest.Config, stats, statsDone)

	for i := 1; i <= loadTest.Config.UniqueIds; i++ {
		context := new(Context)
		context.Config = loadTest.Config
		context.Values = make(map[string]interface{})
		context.CurrentId = i
		context.Stats = stats
		go scenary.Run(context)
	}

	loadTest.Config.wg.Wait()

	close(stats)

	<-statsDone
	return nil
}

// TODO: Essa rotina poderia apenas coletar, e ter outra para gerar os dados para tela, CSV, ...
func collectStats(config *Config, stats chan Stats, done chan bool) {
	startTime := time.Now()

	// TODO: Hoje só temos um cenário. Mas a rotina deve ser revista para trabalhar com mais de um cenário.
	numberOfScenarios := 0
	scenariosWithError := 0
	var minTime time.Duration = 1<<63 - 1
	var maxTime time.Duration
	var totalTime time.Duration
	var scenarioName string

	for elem := range stats {
		fmt.Printf("elem.EndpointID: [%v]\n", elem.EndpointID)
		if elem.MustStat == false {
			continue
		}
		if elem.EndpointID != "" {
			continue
		}
		numberOfScenarios++
		scenarioName = elem.ScenarioID
		log.Println(elem)
		if elem.Status == false {
			scenariosWithError++
		}
		totalTime += elem.Duration
		if elem.Duration > maxTime {
			maxTime = elem.Duration
		}
		if elem.Duration < minTime {
			minTime = elem.Duration
		}
	}

	duration := time.Since(startTime)

	// TODO: Não está sendo separado entre cenários com e sem erros
	log.Printf("Report - Geral")
	log.Printf("\tTempo total de execução do teste: %v", duration)
	log.Printf("\tNúmero de IDs únicos: %v", config.UniqueIds)
	log.Printf("\tCenários executados sem erros: %v (%.2f%%)", numberOfScenarios-scenariosWithError, float64(numberOfScenarios-scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\tCenários executados com erros: %v (%.2f%%)", scenariosWithError, float64(scenariosWithError)/float64(numberOfScenarios)*100.0)
	// log.Printf("\tNúmero de cenários OK: %v", 1)
	// log.Printf("\tNúmero de cenários com falhas: %v", 1)
	log.Printf("\tNúmero de vezes que ocorreu timeout: %v", 0)
	// log.Printf("\tTempos de resposta: Mínimo, médio, máximo, percentil 95%%: %v", 1)

	log.Printf("Report - Por cenário")

	// TODO: tem que fazer o report de todos os cenarios... hoje ta assumindo que so tem 1
	log.Printf("\tCenário: %v", scenarioName)
	log.Printf("\t\tTempo total de execução do cenário: %v", totalTime)
	// log.Printf("\t\tCenários executados: %v", numberOfScenarios) // ou nro de execuções?
	log.Printf("\t\tCenários executados sem erros: %v (%.2f%%)", numberOfScenarios-scenariosWithError, float64(numberOfScenarios-scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\t\tCenários executados com erros: %v (%.2f%%)", scenariosWithError, float64(scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\t\tNúmero de vezes que ocorreu timeout: %v", 1)
	// TODO: faltou percentil 95%
	log.Printf("\t\tTempo de execução (min/med/max): (%v/%v/%v)", minTime, totalTime.Nanoseconds()/int64(numberOfScenarios), maxTime)

	log.Printf("Report - Por endpoint")
	log.Printf("\tEndpoint: %v", "xxx")
	log.Printf("\t\tTempo total de execução do endpoint: %v", 1)
	log.Printf("\t\tEndpoints executados sem erros: %v (%.2f%%)", numberOfScenarios-scenariosWithError, float64(numberOfScenarios-scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\t\tEndpoints executados com erros: %v (%.2f%%)", scenariosWithError, float64(scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\t\tNúmero de vezes que ocorreu timeout: %v", 1)
	// TODO: faltou percentil 95%
	log.Printf("\t\tTempo de execução (min/med/max): (%v/%v/%v)", minTime, totalTime.Nanoseconds()/int64(numberOfScenarios), maxTime)
	// * Tamanho das requisições/respostas: Mínimo, médio, máximo, percentil 95% (percentil tb???)

	// TODO: Report por cenario/endpoint?
	done <- true
}
