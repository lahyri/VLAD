package command

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//VampireRoll - Generates n d10 rolls and compares them to the difficulty
func VampireRoll(n, d int, s *discordgo.Session, m *discordgo.MessageCreate) {
	result := "Rolled: "
	success := 0
	critFailure := false
	had10 := false
	for i := 0; i < n; i++ {
		val := rand.Intn(10) + 1
		result += strconv.Itoa(val) + " "
		if val >= d {
			success++
			if val == 10 {
				had10 = true
			}
		} else if val == 1 {
			critFailure = true
			success--
		}

	}
	s.ChannelMessageSend(m.ChannelID, result)
	if critFailure && (success <= 0) && !(had10) {
		s.ChannelMessageSend(m.ChannelID, "Critical Failure")
	} else if success > 0 {
		response := "You've had " + strconv.Itoa(success) + ("successes")
		s.ChannelMessageSend(m.ChannelID, response)
	} else {
		s.ChannelMessageSend(m.ChannelID, "You've failed.")
	}
}

//DnDRoll - Rolls n d-sided dices, sums them all with b bonus
func DnDRoll(n, d, b int, s *discordgo.Session, m *discordgo.MessageCreate) {
	result := "Rolled: "
	total := b
	for i := 0; i < n; i++ {
		val := rand.Intn(d) + 1
		result += strconv.Itoa(val) + " "
		total += val
	}
	result += strconv.Itoa(b)
	result += " = " + strconv.Itoa(total)
	s.ChannelMessageSend(m.ChannelID, result)
}
