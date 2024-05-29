package command

import "github.com/bwmarrin/discordgo"

var JoinCommand = &discordgo.ApplicationCommand{
	Name:        "join",
	Description: "Start tts in the voice channel",
	DescriptionLocalizations: &map[discordgo.Locale]string{
		discordgo.EnglishGB: "Start tts in the voice channel",
		discordgo.Japanese:  "読み上げを開始します",
	},
	Type: discordgo.ChatApplicationCommand,
}
