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
		case "r":
			n, err := strconv.Atoi(fullCommand[1])
			d := 0
			if len(fullCommand) >= 3 {
				d, _ = strconv.Atoi(fullCommand[2])
			}

			if err == nil {
				command.Roll(n, d, s, m)
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
