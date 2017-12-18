# storm
Storm é uma ferramenta que realiza teste de carga/performance em API's RESTful

Configuração
* URL base: http://localhost:8080/api/
* Timeout máximo é de 60s
* Ids únicos são criados em $test_id e variam entre 1 e 1000
* Slaves é 10
* Threads é 20000
* Processos é 10
* Testes em paralelo é 100
* Incremento de IDs é 10 por segundo
* Respeitar sleep (s/n)
* Constantes?


Cenário: Jornada do usuário com múltiplos usuários
* Dado que eu quero executar o login do usuário
* E o endpoint login é /usuarios/login
* E o JSON é {“email”: “usuario{{$test_id}}@teste_performance.com”, “senha”: “abc123”}
* E o tempo de resposta é de até 3s
* Quando eu envio o POST
* Então eu recebo o status 200
* E aguardo entre 5 e 10 segundos
* E eu guardo a resposta em $login

* Dado que eu quero criar um contato
* E o endpoint cadastro de contatos é /usuarios/{{$login.id_usuario}}/contatos
* E o JSON é {“nome”: “nome do contato”}
* E o tempo de resposta é de até 3s
* Quando eu envio o POST
* Então eu recebo o status 201
* E aguardo entre 5 e 10 segundos
* E eu guardo a resposta em $contato


Report - Geral
* Tempo de execução
* Número de cenários com falhas
* Número de cenários OK
* Número de IDs únicos
* Número de warnings (nro vezes excedeu o timeout)
* Tempos de resposta: Mínimo, máximo, médio, percentil 95%
* Tamanho das requisições/respostas: Mínimo, máximo, médio, percentil 95% (percentil tb???)

Report - Por cenário
* Tempo de execução
* Número de cenários com falhas
* Número de cenários OK
* Número de warnings (nro vezes excedeu o timeout)
* Tempos de resposta: Mínimo, máximo, médio, percentil 95%
* Tamanho das requisições/respostas: Mínimo, máximo, médio, percentil 95% (percentil tb???)

Report - Por endpoint
* Tempo de execução
* Número de cenários com falhas
* Número de cenários OK
* Número de warnings (nro vezes excedeu o timeout)
* Tempos de resposta: Mínimo, máximo, médio, percentil 95%
* Tamanho das requisições/respostas: Mínimo, máximo, médio, percentil 95% (percentil tb???)

TODO
* Com ramp up deve ter os reports acima conforme o número de IDs úncos
  * Com 1, …
  * Com 10, …
  * Com 100, …
* Conforme a carga (usuários), mostrar tempo de resposta, erros, bytes por segundo, …
* Quando pertinente, alem de motrar o TS atual, mostrar o tempo desde que o teste iniciou
* Testes por tempo de execucao ou por bateria (executa todo o teste e para)
* Todos os reports acima podem ser feitos comparando 2 branchs distintos


COMO VAI SER FEITO
* Golang - Lib para carregar formato Gherkin
* Distribuída entre várias máquinas
* Docker?
* Reports no console, S3, slack, e-mail, …
* Histórico de execução no InfluxDB
  * Comparar com execução passada

PLANOS
* Rodar no pipeline para avaliar o desempenho da nova versão

LIBS
* gherkin: https://golanglibs.com/top?q=gherkin
