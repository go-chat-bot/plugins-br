package cnpj

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestCNPJWhenItPassedValidCNPJForValidation(t *testing.T) {
	validCNPJ := "99999999000191"

	bot := getCommandCNPJ()
	bot.Args = []string{validCNPJ}
	got, err := cnpj(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := fmt.Sprintf(msgFmtCnpjValido, validCNPJ)
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCNPJWhenItPassedInvalidCNPJForValidation(t *testing.T) {
	invalidCNPJ := "99999999000100"

	bot := getCommandCNPJ()
	bot.Args = []string{invalidCNPJ}
	got, err := cnpj(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := fmt.Sprintf(msgFmtCnpjInvalido, invalidCNPJ)
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCNPJWhenNoParameterPassedMustGenerateOnlyOneCNPJ(t *testing.T) {
	bot := getCommandCNPJ()
	got, err := cnpj(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	amount := amountOfGeneratedCNPJ(got)
	if amount != 1 {
		t.Errorf("Should return only 1 CNPJ, but got '%d' instead", amount)
	}

	if !valid(strings.Trim(got, " ")) {
		t.Errorf("The generated CNPJ should be valid.")
	}
}

func TestCNPJWhenPassedAQuantityOfCNPJToGenerate(t *testing.T) {
	bot := getCommandCNPJ()
	bot.Args = []string{"3"}

	got, err := cnpj(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	amount := amountOfGeneratedCNPJ(got)
	if amount != 3 {
		t.Errorf("Should return 3 CNPJ, but got '%d' instead", amount)
	}
}

func TestCNPJWhenPassedInvalidParameter(t *testing.T) {
	invalidParameter := "123"
	bot := getCommandCNPJ()
	bot.Args = []string{invalidParameter}

	got, err := cnpj(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := fmt.Sprintf(msgFmtCnpjInvalido, invalidParameter)
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCNPJWhenPassedCNPJWithRepeteadNumbersMustInvalidate(t *testing.T) {
	bot := getCommandCNPJ()
	for i := 0; i <= 9; i++ {
		invalidCNPJ := strings.Repeat(string(i), 14)

		bot.Args = []string{invalidCNPJ}
		got, err := cnpj(bot)

		if err != nil {
			t.Errorf("Error should be nil => %s", err)
		}

		expected := fmt.Sprintf(msgFmtCnpjInvalido, invalidCNPJ)
		if got != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
		}
	}
}

func getCommandCNPJ() *bot.Cmd {
	return &bot.Cmd{
		Command: "cnpj",
	}
}

func amountOfGeneratedCNPJ(r string) int {
	return len(strings.Split(strings.Trim(r, " "), " "))
}
