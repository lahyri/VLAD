package command

import (
	"bufio"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

//help - displays the command list
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	path, _ := os.Getwd()
	path += "/info/help.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.ChannelMessageSend(m.ChannelID, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
