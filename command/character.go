package command

import (
	"ioutil"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

func NewVampire (name string, s *discordgo.Session, m *discordgo.MessageCreate){
	vamp := character.VampireCharacter{
		Name: name
	}
	filename := name + ".json"
	character, err := json.Marshal(vamp)
	err = ioutil.WriteFile(filename, character, 0777)
}