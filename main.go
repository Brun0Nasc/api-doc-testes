package main

import (
	"log"
	"net/http"

	"www.github.com/brun0nasc/api-doc-testes/server"
)

func main(){
	tratador := http.HandlerFunc(server.ServidorJogador)
	if err := http.ListenAndServe(":3030", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 3030 %v", err)
	}
}