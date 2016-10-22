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
			"USD":{"idreg":"215806","code":"USD","codein":"BRL","name":"D\u00f3lar Comercial","high":"3.1676","pctChange":"0.688","open":"0","bid":"3.16","ask":"3.1606","timestamp":"1477076400000","low":"3.1456","notFresh":"1","varBid":"0.0216","create_date":"2016-10-22 13:10:10"},
			"USDT":{"idreg":"215807","code":"USD","codein":"BRLT","name":"D\u00f3lar Turismo","high":"3.15","pctChange":"0.302","open":"0","bid":"3","ask":"3.32","timestamp":"1477076400000","low":"2.99","notFresh":"1","varBid":"0.01","create_date":"2016-10-22 13:10:19"},
			"CAD":{"idreg":"215808","code":"CAD","codein":"BRL","name":"D\u00f3lar Canadense (R$)","high":"2.3862","pctChange":"-0.311","open":"2","bid":"2.3685","ask":"2.3707","timestamp":"1477076340000","low":"2.3641","notFresh":"1","varBid":"-0.0074","create_date":"2016-10-22 13:10:29"},
			"EUR":{"idreg":"215809","code":"EUR","codein":"BRL","name":"Euro (R$)","high":"3.4433","pctChange":"0.049","open":"3","bid":"3.4362","ask":"3.4387","timestamp":"1477076340000","low":"3.4109","notFresh":"1","varBid":"0.0017","create_date":"2016-10-22 13:10:38"},
			"GBP":{"idreg":"215810","code":"GBP","codein":"BRL","name":"Libra Esterlina (R$)","high":"3.8742","pctChange":"0.332","open":"4","bid":"3.8634","ask":"3.8668","timestamp":"1477076400000","low":"3.8273","notFresh":"1","varBid":"0.0128","create_date":"2016-10-22 13:10:48"},
			"ARS":{"idreg":"215811","code":"ARS","codein":"BRL","name":"Peso Argentino (R$)","high":"0.2093","pctChange":"0.82","open":"0","bid":"0.2088","ask":"0.209","timestamp":"1477076340000","low":"0.2064","notFresh":"1","varBid":"0.0017","create_date":"2016-10-22 13:10:57"}
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
	expected := "Dólar: 3.1606 (0.0216), Euro: 3.4387 (0.0017)"
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
