package main

import (
	"log" //* O pacote "log" é usado para exibir mensagens de erro e depuração.
	"net/http" //* O pacote "net/http" é usado para criar um servidor HTTP e lidar com solicitações e respostas HTTP.
	"github.com/Brun0Nasc/api-doc-testes/server"
)

func main(){
	tratador := http.HandlerFunc(server.ServidorJogador) 
	/*
	*A função http.HandlerFunc() é usada para converter
	*a função server.ServidorJogador() em um tipo http.Handler 
	*que pode ser usado para lidar com solicitações HTTP.
	*/
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}