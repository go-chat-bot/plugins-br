package lula

import (
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestLulaWhenTheTextDoesNotMatchLula(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "My name is go-bot, I am awesome."
	got, err := lula(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != "" {
		t.Errorf("Test failed. Expected a empty return, got:  '%s'", got)
	}
}

func TestLulaWhenTheTextMatchLula(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "eu não votei na lula!"
	got, err := lula(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != resp {
		t.Errorf("Test failed. Should return %s", resp)
	}
}
