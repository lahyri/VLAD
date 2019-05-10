package handler

import (
	"strconv"
	"strings"

	"github.com/lahyri/VLAD/command"

	"github.com/bwmarrin/discordgo"
)

// CommandHandler - This function will be called (due to AddHandler in main.go) every time
// a new message is created on any channel that the autenticated bot has access to.
func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
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
			d := 0
			if len(fullCommand) >= 3 {
				d, _ = strconv.Atoi(fullCommand[2])
			}

			if err1 == nil {
				command.VampireRoll(n, d, s, m)
			}

		case "dndr":
			n, err1 := strconv.Atoi(fullCommand[1])
			d, err2 := strconv.Atoi(fullCommand[2])
			b := 0
			if len(fullCommand) >= 4 {
				b, _ = strconv.Atoi(fullCommand[3])
			}

			if err1 == nil && err2 == nil {
				command.DnDRoll(n, d, b, s, m)
			}
		case "help":
			command.Help(s, m)

		case "discipline":
			l := 0
			d := "disciplines"
			if len(fullCommand) >= 2 {
				d = fullCommand[1]
				if len(fullCommand) >= 3 {
					l, _ = strconv.Atoi(fullCommand[2])
				}
			}

			command.Discipline(d, l, s, m)
		}

	}

}
