package readsense

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Format is a function that formats the message.
type Options struct {
	ReadUsername     bool
	MaxMessageLength int
}

func Format(message *discordgo.Message) string {
	base := message.Content

	// Replace line breaks with space
	base = strings.ReplaceAll(base, "\n", " ")

	// Replace mentions with username
	for _, m := range message.Mentions {
		base = strings.ReplaceAll(base, "<@"+m.ID+">", m.Username)
	}

	// Replace mentions with role
	for _, r := range message.MentionRoles {
		base = strings.ReplaceAll(base, "<@&"+r+">", r)
	}

	// Replace custom emojis
	for _, e := range message.GetCustomEmojis() {
		base = strings.ReplaceAll(base, "<:"+e.Name+":"+e.ID+">", e.Name)
	}

	// Replace attachments
	for _, a := range message.Attachments {
		base = strings.ReplaceAll(base, a.URL, "ゆーあーるえる")
	}

	numsAttachments := len(message.Attachments)
	if numsAttachments > 0 {
		base = "ファイル添付。" + base
	}

	return base
}
