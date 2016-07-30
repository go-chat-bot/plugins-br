package megasena

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/go-chat-bot/bot"
)

const (
	retornoJSON = `{"concurso":{
    "numero":"1636",
    "data":"17\/09\/2014",
    "cidade":"OSASCO-SP",
    "local":"Caminh\u00e3o da Sorte",
    "valor_acumulado":"29.530.043,53",
    "numeros_sorteados":[
      "19",
      "26",
      "33",
      "35",
      "51",
      "52"
    ],
    "premiacao":{
      "sena":{
        "ganhadores":"0",
        "valor_pago":"0,00"
      },
      "quina":{
        "ganhadores":"90",
        "valor_pago":"38.637,27"
      },
      "quadra":{
        "ganhadores":"8.474",
        "valor_pago":"586,22"
      }
    },
    "arrecadacao_total":"59.395.800,00"
  },
  "proximo_concurso":{
    "data":"20\/09\/2014",
    "valor_estimado":"37.000.000,00"
  },
  "concurso_final_zero":{
    "numero":"1640",
    "valor_acumulado":"7.271.924,76"
  },
  "mega_virada_valor_acumulado":"54.516.366,32"}`
)

func TestSortear(t *testing.T) {
	esperado := "01 02 03 04 05 06"
	resultado := sortear(6)
	if resultado != esperado {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", esperado, resultado)
	}
}

func getCommandMegasena() *bot.Cmd {
	return &bot.Cmd{
		Command: "megasena",
	}
}

func TestMegaSenaQuandoNaoEPassadoArgumento(t *testing.T) {
	cmd := getCommandMegasena()
	cmd.Args = []string{}
	got, err := megasena(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != msgOpcaoInvalida {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", msgOpcaoInvalida, got)
	}
}

func TestMegaSenaQuandoArgumentoForGerar(t *testing.T) {
	cmd := getCommandMegasena()
	cmd.Args = []string{"gerar"}
	got, err := megasena(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	match, err := regexp.MatchString("(\\d{2} {1}){5}\\d{2}", got)
	if err != nil {
		t.Errorf("Failed match: %s", err)
	}
	if !match {
		t.Errorf("Test failed, match should be true")
	}
}

func TestMegaSenaQuandoArgumentoForResultado(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, retornoJSON)
		}))
	defer ts.Close()

	url = ts.URL

	cmd := getCommandMegasena()
	cmd.Args = []string{"resultado"}
	got, err := megasena(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := "Sorteio 1636 de 17/09/2014: [19 26 33 35 51 52] - 0 premiado(s) R$ 0,00."
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestMegaSenaQuandoArgumentoForResultadoERetornoInvalido(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "invalid")
		}))
	defer ts.Close()

	url = ts.URL

	cmd := getCommandMegasena()
	cmd.Args = []string{"resultado"}
	_, err := megasena(cmd)

	if err == nil {
		t.Errorf("Error shouldn't be nil")
	}
}
