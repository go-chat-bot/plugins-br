package gloria_a_deus

import (
	"strings"
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestGloriaADeusWhenTheTextDoesNotMatchGloriaADeus(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "My name is go-bot, I am awesome."
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != "" {
		t.Errorf("Test failed. Expected a empty return, got:  '%s'", got)
	}
}

func TestGloriaADeusWhenTtheTextMatchGloriaADeus(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Gloria A Deus!"
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Gloria a Deuxxx!") {
		t.Errorf("Test failed. Should return a clever Gloria A Deux quote")
	}
}

func TestGloriaADeusWhenTtheTextMatchJesus(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Jesus!"
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Gloria a Deuxxx!") {
		t.Errorf("Test failed. Should return a clever Gloria A Deux quote")
	}
}

func TestGloriaADeusWhenTtheTextMatchGod(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Oh my God!"
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Gloria a Deuxxx!") {
		t.Errorf("Test failed. Should return a clever Gloria A Deux quote")
	}
}

func TestGloriaADeusWhenTtheTextMatchGloria(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Oh Gloria!"
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Gloria a Deuxxx!") {
		t.Errorf("Test failed. Should return a clever Gloria A Deux quote")
	}
}

func TestGloriaADeusWhenTtheTextMatchGloria1(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Oh Glória!"
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Gloria a Deuxxx!") {
		t.Errorf("Test failed. Should return a clever Gloria A Deux quote")
	}
}

func TestGloriaADeusWhenTtheTextMatchDaciolo(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Foi o Daciolo!"
	got, err := gloria_a_deus(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Gloria a Deuxxx!") {
		t.Errorf("Test failed. Should return a clever Gloria A Deux quote")
	}
}
