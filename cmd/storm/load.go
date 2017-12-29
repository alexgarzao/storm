package main

import (
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
	// Fluxo principal
	//
	// Carrega config
	// Inicializa dados a serem medidos
	// Configura número processos
	// Dispara threads conforme número usuários
	// 	Inicializa medições cenário
	// 	Para cada step
	// 		Inicializa medições step
	// 		Executa step
	// 		Se erro aborta
	// 		Coleta dados step
	// 	Coleta dados cenário
	// Coleta dados teste
	// Gera report

	loadTest.Config.wg.Add(loadTest.Config.UniqueIds)

	stats := make(chan Stats, loadTest.Config.UniqueIds)

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

	printStats(stats)

	return nil
}

func printStats(stats chan Stats) {
	// Estatísticas por cenário:
	//     Rótulo (nome do cenário)
	//     Número cenários executados (com erro / corretas) - nro e %
	//     Tempo mínimo/médio/máximo
	//     Bytes (média) retornados pelo servidor, ...

	// TODO: Hoje só temos um cenário. Mas a rotina deve ser revista para trabalhar com mais de um cenário.
	numberOfScenarios := len(stats)
	scenariosWithError := 0
	var minTime time.Duration = 999999
	var maxTime time.Duration
	var totalTime time.Duration
	var scenarioName string

	for elem := range stats {
		scenarioName = elem.Id
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

	// TODO: Não está sendo separado entre cenários com e sem erros
	log.Printf("Cenário %s executou em %v", scenarioName, totalTime)
	log.Printf("\tCenários executados: %v", numberOfScenarios)
	log.Printf("\tCenários executados sem erros: %v (%.2f%%)", numberOfScenarios-scenariosWithError, float64(numberOfScenarios-scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\tCenários executados com erros: %v (%.2f%%)", scenariosWithError, float64(scenariosWithError)/float64(numberOfScenarios)*100.0)
	log.Printf("\tTempo de execução (min/med/max): (%v/%v/%v)", minTime, totalTime.Nanoseconds()/int64(numberOfScenarios), maxTime)
}
