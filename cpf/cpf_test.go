package cpf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestCPFWhenItPassedValidCPFForValidation(t *testing.T) {
	validCPF := "52998224725"

	bot := getCommandCPF()
	bot.Args = []string{validCPF}

	got, err := cpf(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := fmt.Sprintf(msgFmtCpfValido, validCPF)
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCPFWhenItPassedInvalidCPFForValidation(t *testing.T) {
	invalidCPF := "52998224700"

	bot := getCommandCPF()
	bot.Args = []string{invalidCPF}

	got, err := cpf(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := fmt.Sprintf(msgFmtCpfInvalido, invalidCPF)
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCPFWhenNoParameterPassedMustGenerateOnlyOneCPF(t *testing.T) {
	bot := getCommandCPF()
	got, err := cpf(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	amount := amountOfGeneratedCPF(got)
	if amount != 1 {
		t.Errorf("Should return only 1 CPF, but got '%d' instead", amount)
	}

	if !valid(strings.Trim(got, " ")) {
		t.Errorf("The generated CPF should be valid.")
	}
}

func TestCPFWhenPassedAQuantityOfCPFToGenerate(t *testing.T) {
	bot := getCommandCPF()
	bot.Args = []string{"3"}

	got, err := cpf(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	amount := amountOfGeneratedCPF(got)
	if amount != 3 {
		t.Errorf("Should return 3 CPF, but got '%d' instead", amount)
	}
}

func TestCPFWhenPassedInvalidParameter(t *testing.T) {
	invalidParameter := "123"
	bot := getCommandCPF()
	bot.Args = []string{invalidParameter}

	got, err := cpf(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}

	expected := fmt.Sprintf(msgFmtCpfInvalido, invalidParameter)
	if got != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCPFWhenPassedCPFWithRepeteadNumbersMustInvalidate(t *testing.T) {
	bot := getCommandCPF()
	for i := 0; i <= 9; i++ {
		invalidCPF := strings.Repeat(string(i), 11)

		bot.Args = []string{invalidCPF}
		got, err := cpf(bot)

		if err != nil {
			t.Errorf("Error should be nil => %s", err)
		}

		expected := fmt.Sprintf(msgFmtCpfInvalido, invalidCPF)
		if got != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, got)
		}
	}
}

func getCommandCPF() *bot.Cmd {
	return &bot.Cmd{
		Command: "cpf",
	}
}

func amountOfGeneratedCPF(r string) int {
	return len(strings.Split(strings.Trim(r, " "), " "))
}
