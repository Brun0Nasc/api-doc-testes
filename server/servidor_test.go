package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
*Este código é um teste unitário para o pacote server. 
*Ele testa a função ServeHTTP do ServidorJogador para garantir
*que ela está respondendo corretamente às requisições HTTP.
*/

type EsbocoArmazenamentoJogador struct {
	potuacoes map[string]int
}

/*
*O teste usa uma implementação de esboço do ArmazenamentoJogador 
*chamada EsbocoArmazenamentoJogador, que contém um mapa que mapeia 
*nomes de jogadores a suas pontuações. O esboço é usado como fonte 
*de dados para a função ObterPontuacaoJogador do ServidorJogador.
*/

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.potuacoes[nome]
	return pontuacao
}

func comparaResultados(t *testing.T, recebido, esperado string){
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
	}
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
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

		/*
		*O teste usa a função novaRequisicaoObterPontuacao para criar uma nova requisição 
		*HTTP com o método GET e o caminho correspondente ao jogador desejado. Em seguida, 
		*a função ServeHTTP do ServidorJogador é chamada com essa requisição e um gravador 
		*de resposta httptest.NewRecorder.
		*/

		servidor.ServeHTTP(resposta, requisicao)

		comparaResultados(t, resposta.Body.String(), "20")

		/*
		*O teste usa a função comparaResultados para comparar o resultado recebido com o 
		*resultado esperado. Se o resultado recebido não for igual ao resultado esperado,
		*o teste falhará e uma mensagem de erro será exibida.
		*/
	})

	t.Run("retornando resultado de Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		comparaResultados(t, resposta.Body.String(), "10")
	})
}
