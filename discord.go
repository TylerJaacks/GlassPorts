package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	dg         *discordgo.Session
)
	
func startDiscord() {

	dg, _ = discordgo.New("Bot " + config.BotToken)
	dg.AddHandler(messageCreate)
	_ = dg.Open()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	LogMsg("Detected incoming message.")
	// Return if message came from a bot, or doesn't mention this bot
	if m.Author.Bot || !strings.Contains(m.Content, s.State.User.ID) || !strings.Contains(config.GuildID, m.GuildID) {
		LogMsg("Ignoring message.")
		return
	}
	// Split input for use in command functions
	parts := strings.Split(m.Content, " ")
	b := BotCommand{
		Session:   s,
		Channel:   m.ChannelID,
		Message:   m,
		Command:   parts[1],
		DiscordID: m.Author.ID,
		Parts:     parts,
	}
	LogMsg("Command detected: %+v", b)

	// No valid command found
	b.Reply(fmt.Sprintf("Unknown command: %+v. Ping me with **iaadd** followed by a number or a link to the in-game screenshot of your score to record your Infinity Arena high score. Ping me with **pvpadd** followed by a number or a link to the in-game screenshot of your leaderboard high score to record your PvP leaderboard score. Ping me with **iacheck** or **pvpcheck** to request the information you have inserted.", b.Command))

}

// Reply will reply to the BotCommand.Message, tagging the sender. If b.Response is set, it will use that otherwise the string will be used
func (b BotCommand) Reply(s string) {
	if len(b.Response) > 0 {
		b.Session.ChannelMessageSend(b.Channel, fmt.Sprintf("<@%+v>: %+v", b.DiscordID, b.Response))
	} else {
		b.Session.ChannelMessageSend(b.Channel, fmt.Sprintf("<@%+v>: %+v", b.DiscordID, s))
	}
}