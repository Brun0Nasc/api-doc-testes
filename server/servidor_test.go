package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func comparaResultados(t *testing.T, recebido, esperado string){
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
	}
}

func TestObterJogadores(t *testing.T) {
	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Maria", nil)
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "20"

		comparaResultados(t, recebido, esperado)
	})

	t.Run("retornando resultado de Pedro", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Pedro", nil)
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "10"

		comparaResultados(t, recebido, esperado)
	})
}