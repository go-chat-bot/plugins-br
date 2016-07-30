package dilma

import (
	"strings"
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestDilmaWhenTheTextDoesNotMatchDilma(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "My name is go-bot, I am awesome."
	got, err := dilma(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != "" {
		t.Errorf("Test failed. Expected a empty return, got:  '%s'", got)
	}
}

func TestDilmaWhenTtheTextMatchDilma(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "eu não votei na dilma!"
	got, err := dilma(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, ":dilma: ") {
		t.Errorf("Test failed. Should return a clever Dilma quote")
	}
}
