package gloria_a_deus

import (
	"fmt"
	"regexp"

	"github.com/go-chat-bot/bot"
)

const (
	pattern = "(?i)\\b(deus|jesus|god|gloria|glória)\\b"
)

var (
	re = regexp.MustCompile(pattern)
)

func gloria_a_deus(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {
		return fmt.Sprintf("Gloria a Deuxxx!"), nil
	}
	return "", nil
}

func init() {
	bot.RegisterPassiveCommand(
		"gloria_a_deus",
		gloria_a_deus)
}
