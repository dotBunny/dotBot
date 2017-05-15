# dotBot
## About
This is just a helper bot that sits in our Discord server, as well as the PTBO Game Jam server. It's nothing too fancy, but something here might be useful to someone, so best to make it public.

## Requirements
You will need some place to run the bot with a functioning Go workspace. Other then that, it is pretty self explanitory.

- Go Lang 1.7.5+

### External Packages
- github.com/BurntSushi/toml
- github.com/bwmarrin/discordgo
- github.com/mmcdole/gofeed
- golang.org/x/net/html

## Installation

Once you've got the repo locally, you simply need to make sure to create a "bin" folder and put a *config.toml* in there. The binary will compile to there as well.

The contents of your *config.toml* should look like this:

> token = "_put your bot token here_"
