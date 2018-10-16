package command

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Detects the "!" to see if the message is a command
	if m.Content[:1] == "!" {
		fullCommand := strings.Split(m.Content, " ")
		method := fullCommand[0][1:]

		switch method {
		case "vr":
			n, err1 := strconv.Atoi(fullCommand[1])
			d, err2 := strconv.Atoi(fullCommand[2])
			if err2 != nil {
				d = 0
			}
			if err1 == nil {
				VampireRoll(n, d, s, m)
			}

		case "dndr":
			n, err1 := strconv.Atoi(fullCommand[1])
			d, err2 := strconv.Atoi(fullCommand[2])
			b, err3 := strconv.Atoi(fullCommand[3])
			if err3 != nil {
				b = 0
			}
			if err1 == nil && err2 == nil {
				DnDRoll(n, d, b, s, m)
			}
		case "help":
		default:
			Help(s, m)
		}

	}

}
