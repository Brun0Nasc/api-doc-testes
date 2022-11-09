package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObterJogadores(t *testing.T) {
	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "20"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}
	})

	t.Run("retornar resultado de Pedro Migule", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "10"

		verificaCorpoRequisicao(t, recebido, esperado)
	})
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificaCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
	}
}