package server

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}

type ServidorJogador struct {
	Armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request){
	jogador := r.URL.Path[len("/jogadores/"):]

	pontuacao := s.Armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}