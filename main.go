package main

import (
	"log" //* O pacote "log" é usado para exibir mensagens de erro e depuração.
	"net/http" //* O pacote "net/http" é usado para criar um servidor HTTP e lidar com solicitações e respostas HTTP.
	"github.com/Brun0Nasc/api-doc-testes/server"
)

type ArmazenamentoJogadorEmMemoria struct{}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return 123
}

func main(){
	servidor := &server.ServidorJogador{Armazenamento:&ArmazenamentoJogadorEmMemoria{}}

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}