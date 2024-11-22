package main

import "github.com/bwmarrin/discordgo"

type Config struct {
	GuildID string

	DatabaseInfo string

	DBUsername string

	DBPassword string

	BotToken string
}

type User struct {
	Identifier string

	Upload string

	Username string

	Email string

	Bio string

	Location string

	Twitter string

	Github string

	Keybase string

	Discord string

	LinkedIn string

	Website string

	Reddit string

	ProfilePic string
}

type GlassApp struct {
	AppName string

	ShortDesc string

	LongDesc string

	APKLink string

	AppID string

	Screenshots string

	Maintainer string

	Icon string
}

type CompanionApp struct {
	AppName string

	APKLink string

	GlassAppID string

	AppID string

	Icon string
}

type DownloadStats struct {
	Identifier string

	Rating string

	Review string

	GlassAppID string
}

type AuthToken struct {
	Identifier string
	Email string
	Username string
	PasswordHash string
	LastIP string
	AuthToken string
	PreviousHash string
	Current string
}

type BotCommand struct {
	Channel   string
	DiscordID string
	Command   string
	Message   *discordgo.MessageCreate
	Session   *discordgo.Session
	Parts     []string
	Response  string
}

type APIResponse struct {
	Code int
	Valid bool
	Response interface{}
}
