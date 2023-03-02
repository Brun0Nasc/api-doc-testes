package server

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request){
	jogador := r.URL.Path[len("/jogadores/"):]
	fmt.Fprint(w, s.armazenamento.ObterPontuacaoJogador(jogador))
}