package command

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Discipline(discipline string, level int, s *discordgo.Session, m *discordgo.MessageCreate) {
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
