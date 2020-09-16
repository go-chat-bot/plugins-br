package gloria_a_deus

import (
	"regexp"

	"github.com/go-chat-bot/bot"
)

const (
	pattern = "(?i)\\b(deus|jesus|god|gloria|glória|daciolo)\\b"
)

var (
	re = regexp.MustCompile(pattern)
)

func gloria_a_deus(command *bot.PassiveCmd) (s string, err error) {
	if re.MatchString(command.Raw) {
		s = "Gloria a Deuxxx!"
	}
	return
}

func init() {
	bot.RegisterPassiveCommand(
		"gloria_a_deus",
		gloria_a_deus)
}
