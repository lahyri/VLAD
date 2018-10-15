package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token          string
	Command        string
	DiceQuantity   int
	DiceDifficulty int
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&Command, "cmd", "help", "specifies the command used")
	flag.IntVar(&DiceQuantity, "dq", 0, "defines the ammount of dices to be rolled")
	flag.IntVar(&DiceDifficulty, "dd", 0, "defines the difficulty of the test")

	flag.Parse()

	if Token == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
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
	// If the message is "ping" reply with "Pong!"
	if Command == "!r" {
		result := ""
		success := 0
		critFailure := false
		for i := 0; i < DiceQuantity; i++ {
			val := rand.Intn(10)
			result += strconv.Itoa(val) + " "
			if val >= DiceDifficulty {
				success++
			} else if val == 1 {
				critFailure = true
				success--
			}
		}
		s.ChannelMessageSend(m.ChannelID, result)
		if critFailure && success <= 0 {
			s.ChannelMessageSend(m.ChannelID, "Critical Failure")
		} else if success > 0 {
			response := "You've had " + strconv.Itoa(success) + ("successes")
			s.ChannelMessageSend(m.ChannelID, response)
		} else {
			s.ChannelMessageSend(m.ChannelID, "You've failed.")
		}

	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
