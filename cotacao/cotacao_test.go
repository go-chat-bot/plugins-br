package cotacao

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chat-bot/bot"
)

const (
	expectedJSON = `{
      "bovespa":{
        "cotacao":"60800",
        "variacao":"-1.68"
      },
      "dolar":{
        "cotacao":"2.2430",
        "variacao":"+0.36"
      },
      "euro":{
        "cotacao":"2.9018",
        "variacao":"-1.21"
      },
      "atualizacao":"04\/09\/14   -18:13"
    }`
)

func TestCotacaoMustRespondWithTheDollarAndEuroCurrencyExchange(t *testing.T) {
	cmd := &bot.Cmd{}

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, expectedJSON)
		}))
	defer ts.Close()
	url = ts.URL

	got, err := cotacao(cmd)
	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	expected := "Dólar: 2.2430 (+0.36), Euro: 2.9018 (-1.21)"
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCotacaoWhenWebServiceReturnsSomethingInvalidMustReturnError(t *testing.T) {
	cmd := &bot.Cmd{}

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "invalid")
		}))
	defer ts.Close()
	url = ts.URL

	_, err := cotacao(cmd)
	if err == nil {
		t.Errorf("Error shouldn't be nil")
	}
}
