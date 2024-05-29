package command

import "github.com/bwmarrin/discordgo"

var VoiceSettingsCommands = &discordgo.ApplicationCommand{
	Name:        "voice-settings",
	Description: "modify voice settings",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name: "model",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name: "model",
				},
			},
		},
	},
}
