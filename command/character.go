package command

type Character struct {
	Name                string
	Player              string
	Chronicle           string
	Nature              string
	Demeanor            string
	Concept             string
	Clan                string
	Generation          int
	Sire                string
	Attributes          Attribute
	Abilities           Ability
	Disciplines         map[string]int
	Backgrounds         map[string]int
	Conscience          int
	SelfControl         int
	Courage             int
	Humanity            int
	WillpowerTotal      int
	WillpowerCurrent    int
	BloodpoolTotal      int
	BloodpoolCurrent    int
	Weakness            string
	ExperienceTotal     int
	ExperienceAvailable int
}
type Attribute struct {
	Strength     int
	Dexterity    int
	Stamina      int
	Charisma     int
	Manipulation int
	Appearence   int
	Perception   int
	Intelligence int
	Wits         int
}
type Ability struct {
	//Talents
	Athletics    int
	Brawl        int
	Dodge        int
	Empathy      int
	Expression   int
	Intimidation int
	Leadership   int
	Streetwise   int
	Subterfuge   int

	//Skills
	AnimalKen   int
	Crafts      int
	Drive       int
	Etiquette   int
	Firearms    int
	Melee       int
	Performance int
	Security    int
	Stealth     int
	Survival    int

	//Knowledges
	Academics     int
	Bureaucracy   int
	Computer      int
	Finance       int
	Investigation int
	Law           int
	Linguistics   int
	Medicine      int
	Occult        int
	Politics      int
	Science       int
}
