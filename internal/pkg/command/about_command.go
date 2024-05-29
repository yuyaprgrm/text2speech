package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yuyaprgrm/text2speech/internal/pkg/core"
)

// AboutCommand is the command for displaying bot version
var AboutCommand = &discordgo.ApplicationCommand{
	Name:        "about",
	Description: "displays bot version",
}

var AboutCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "About",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Version",
							Value:  core.VERSION,
							Inline: true,
						},
						{
							Name:   "Source",
							Value:  "[GitHub](" + core.GITHUB_URL + ")",
							Inline: true,
						},
						{
							Name:   "Developer",
							Value:  "yuyaprgrm",
							Inline: true,
						},
						{
							Name:   "License",
							Value:  "MIT",
							Inline: true,
						},
					},
				},
			},
		},
	})
}
