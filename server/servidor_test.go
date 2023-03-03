package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	potuacoes map[string]int
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.potuacoes[nome]
	return pontuacao
}

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Maria":20,
			"Pedro":10,
		},
	}
	servidor := &ServidorJogador{&armazenamento}

	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		comparaResultados(t, resposta.Body.String(), "20")

	})

	t.Run("retornando resultado de Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		comparaResultados(t, resposta.Body.String(), "10")
	})

	t.Run("retorna 404 para o jogador não encontrado", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Jorge")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Code
		esperado := http.StatusNotFound

		if recebido != esperado {
			t.Errorf("recebido status %d esperado %d", recebido, esperado)
		}
	})
}

func comparaResultados(t *testing.T, recebido, esperado string){
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
	}
}

func verificarRespostaCodigoStatus(t *testing.T, recebido, esperado int) {
	t.Helper()

	if recebido != esperado {
		t.Errorf("não recebeu o código de status HTTP esperado, recebido %d, esperado %d", recebido, esperado)
	}
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}
