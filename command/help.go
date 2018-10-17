package command

import (
	"github.com/bwmarrin/discordgo"
)

const (
	BotDescr         = ""
	VampireRollDescr = "blabla"
	DnDRollDescr     = "loblob"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, VampireRollDescr)

}
