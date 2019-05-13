package command

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//Roll - Generates n d10 rolls and compares them to the difficulty
func Roll(n, d int, s *discordgo.Session, m *discordgo.MessageCreate) {
	result := "Rolled: ["
	success := 0
	critFailure := false
	had10 := false
	for i := 0; i < n; i++ {
		val := rand.Intn(10) + 1
		if i == n-1 {
			result += strconv.Itoa(val)
		} else {
			result += strconv.Itoa(val) + ", "
		}
		if val >= d {
			success++
			if val == 10 {
				had10 = true
			}
		}
		if val == 1 {
			critFailure = true
			success--
		}

	}
	result += "]"
	s.ChannelMessageSend(m.ChannelID, result)
	if critFailure && (success <= 0) && !(had10) {
		s.ChannelMessageSend(m.ChannelID, "Critical Failure")
	} else if success > 0 {
		response := "You've had " + strconv.Itoa(success)
		if success == 1 {
			response += " success"
		} else {
			response += " successes"
		}
		s.ChannelMessageSend(m.ChannelID, response)
	} else {
		s.ChannelMessageSend(m.ChannelID, "You've failed.")
	}
}
