package command

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//discipline - Sends the description of the determinated discipline in the specified level. If called without parameters, returns a list with all the disciplines
func discipline(discipline string, level int, s *discordgo.Session, m *discordgo.MessageCreate) {
	path, _ := os.Getwd()
	path += "/info/"
	if discipline == "disciplines" {
		path += "disciplines.txt"
	} else {
		path += discipline + "/" + strconv.Itoa(level) + ".txt"
	}

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
