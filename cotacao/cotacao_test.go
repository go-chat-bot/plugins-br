package cotacao

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chat-bot/bot"
)

const (
	expectedJSON = `
		{"base":"BRL",
		"date":"2017-03-15",
		"rates":{"AUD":0.41701,
		"BGN":0.5829,
		"CAD":0.42592,
		"CHF":0.31935,
		"CNY":2.1887,
		"CZK":8.0532,
		"DKK":2.2156,
		"GBP":0.25951,
		"HKD":2.4596,
		"HRK":2.2143,
		"HUF":92.639,
		"IDR":4229.4,
		"ILS":1.1579,
		"INR":20.787,
		"JPY":36.292,
		"KRW":362.8,
		"MXN":6.1994,
		"MYR":1.4085,
		"NOK":2.7226,
		"NZD":0.45531,
		"PHP":15.912,
		"PLN":1.2877,
		"RON":1.3538,
		"RUB":18.7,
		"SEK":2.8493,
		"SGD":0.44708,
		"THB":11.162,
		"TRY":1.1772,
		"USD":0.31657,
		"ZAR":4.1323,
		"EUR":0.29804}}`
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
	expected := "Dólar: 3.16, Euro: 3.36, CAD: 2.35, Libra: 3.85"
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
