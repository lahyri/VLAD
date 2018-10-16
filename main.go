package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/lahyri/VLAD/command"
)

// Variables used for command line parameters
var (
	Command        string
	DiceQuantity   int
	DiceDifficulty int
)

//Token - Your Discord token
const Token = "Bot INSERT KEY HERE"

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

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
			n, _ := strconv.Atoi(fullCommand[1])
			d, err := strconv.Atoi(fullCommand[2])
			if err != nil {
				d = 0
			}
			command.VampireRoll(n, d, s, m)
		case "help":
		default:
			command.Help()
		}

	}

}
