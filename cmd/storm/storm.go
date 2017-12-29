package main

import (
	"log"
	"runtime"
	"sync"
)

type Config struct {
	BaseUrl           string
	RequestMaxTimeout int
	UniqueIds         int
	wg                sync.WaitGroup
}

func main() {
	log.Printf("Starting Storm V 0.1\n")
	log.Printf("Número de processos: %v", runtime.GOMAXPROCS(0))
	config := &Config{
		BaseUrl:           "http://localhost:8000/api/",
		RequestMaxTimeout: 60 * 1000,
		UniqueIds:         100,
	}
	load := NewLoadTest(config)

	// Cenário: Jornada com múltiplos usuários
	scenary := NewScenary("Jornada com múltiplos usuários")

	// Dado que eu quero executar o login do usuário

	// E o endpoint login é /usuarios/login
	scenary.AddStep(NewStepDefineEndpoint("usuarios/login/"))

	// E o JSON é {“email”: “usuario{{$test_id}}@teste_performance.com”, “senha”: “abc123”}
	scenary.AddStep(NewStepDefineJson("{\"email\": \"usuario{{.CurrentId}}@perform.com\", \"senha\": \"senha_do_usuario{{.CurrentId}}@perform.com\"}"))

	// E o tempo de resposta é de até 3s
	scenary.AddStep(NewStepDefineRequestTimeout(3 * 1000))

	// Quando eu envio o POST
	scenary.AddStep(NewStepSendPost())

	// Então eu recebo o status 200
	scenary.AddStep(NewStepCheckHttpStatusCode(200))

	// E aguardo entre 5 e 10 segundos

	// E eu guardo a resposta em $login

	if err := load.Run(scenary); err != nil {
	}

	log.Printf("Finishing Storm V 0.1\n")
}

// - Inicializa dados a serem medidos
// - Configura número processos
// - Dispara threads conforme número usuários
//     - Inicializa medições cenário
//     - Para cada step
//         - Inicializa medições step
//         - Executa step
//         - Se erro aborta
//         - Coleta dados step
//     - Coleta dados cenário
// - Coleta dados teste
// - Gera report
