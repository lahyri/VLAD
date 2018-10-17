# VLAD
### Vampire Automated Log for Discord

## Installing

To install VLAD in your Discord Server you'll need the following:

* [Golang](https://golang.org/)
* [Discord](https://discordapp.com/) (duuh)

After that, you should run the following command:
```
$go get github.com/bwmarrin/discordgo
```
And your own version of VLAD is ready to roll.

### Observation

Since I'm not planning to host this permanently, You'll have to [create your own bot account](https://discordpy.readthedocs.io/en/rewrite/discord.html), save the generated token to a "token.txt" file in the main project folder an run the project everytime you'll want to use it.

## Running

Open your code folder in the terminal and run:
```
$go run main.go
```
If you've done this correctly, you bot should appear online in the server.

## Commands

To see all the commands available for this bot, just type the message

    !help

in your discord server, and it shall return all the commands and their description of usage.

## Contribuiting

Feel free to use it and modify it as you want. If you'd like to share any new functionality, just branch the rep and send me a pull request with a description of the alteration or functionality you've added.

## To Do

* Disciplines/Skills/Attributes list
* Discipline/Skill/Attribute description
* More stuff I still have to think about