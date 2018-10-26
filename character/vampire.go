package character

type VampireCharacter struct {
	Name                string
	Player              string
	Chronicle           string
	Nature              string
	Demeanor            string
	Concept             string
	Clan                string
	Generation          int
	Sire                string
	Attributes          VAttribute
	Abilities           VAbility
	Disciplines         VDiscipline
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

type VDiscipline struct {
	Animalism     int
	Auspex        int
	Celerity      int
	Chimerstry    int
	Dementation   int
	Dominate      int
	Fortitude     int
	Obfuscate     int
	Obtenebration int
	Potence       int
	Presence      int
	Protean       int
	Quietus       int
	Serpentis     int
	Thanatosis    int
	Vicissitude   int
}

type VAttribute struct {
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
type VAbility struct {
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
