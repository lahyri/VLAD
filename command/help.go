package command

import "github.com/bwmarrin/discordgo"

const (
	BotDescr         = ""
	VampireRollDescr = ""
	DnDRollDescr     = ""
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, BotDescr)
	s.ChannelMessageSend(m.ChannelID, VampireRollDescr)
	s.ChannelMessageSend(m.ChannelID, DnDRollDescr)
}
