package text2speech

import (
	"log/slog"
	"os"
	"path"

	"github.com/bwmarrin/discordgo"
	"github.com/yuyaprgrm/text2speech/internal/pkg/command"
	"github.com/yuyaprgrm/text2speech/internal/pkg/guild"
	"github.com/yuyaprgrm/text2speech/internal/pkg/readsense"
	"github.com/yuyaprgrm/text2speech/internal/pkg/tts"
	"github.com/yuyaprgrm/text2speech/pkg/openjtalk"
)

type Text2Speech struct {
	primary *discordgo.Session
	guilds  map[string]*guild.Guild
}

func NewText2Speech(primary *discordgo.Session) *Text2Speech {
	t := &Text2Speech{
		primary: primary,
	}

	primary.AddHandler(t.ready)
	primary.AddHandler(t.interactionCreate)
	primary.AddHandler(t.guildCreate)
	return t
}

func (t *Text2Speech) Open() {
	t.primary.Open()
}

func (t *Text2Speech) ready(s *discordgo.Session, r *discordgo.Ready) {
	slog.Info("Bot is ready", "user", r.User.Username)
}

func (t *Text2Speech) guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	slog.Info("Joined guild", "guild", g.Guild.ID)
	t.registerCommands(s, g.Guild.ID)
}

func (t *Text2Speech) registerCommands(s *discordgo.Session, guildID string) {
	commands := []*discordgo.ApplicationCommand{
		command.JoinCommand,
		command.AboutCommand,
	}

	slog.Debug("Registering commands", "guild", guildID)
	for _, v := range commands {
		s.ApplicationCommandCreate(s.State.User.ID, guildID, v)
	}
}

func (t *Text2Speech) interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "about":
			command.AboutCommandHandler(s, i)
		case "join":
			slog.Info("finding voice state", "user", i.Member.User.ID)
			userVoiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "You are not in a voice channel",
					},
				})
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "読み上げを開始します",
				},
			})
			startTts(s, i.GuildID, i.ChannelID, userVoiceState.ChannelID)
		}
	}
}

func startTts(s *discordgo.Session, guildID, readChannelID, voiceChannelID string) {
	_, connected := s.VoiceConnections[guildID]
	if connected {
		return
	}
	vc, err := s.ChannelVoiceJoin(guildID, voiceChannelID, false, true)
	if err != nil {
		slog.Error("failed to join voice channel", "error", err)
		return
	}

	engine := tts.NewEngine(
		"/usr/bin/open_jtalk",
		openjtalk.NewCtx().WithDictionaryDirectory("/var/lib/mecab/dic/open-jtalk/naist-jdic"),
	)

	outdir := path.Join("out", voiceChannelID)
	os.MkdirAll(outdir, 0755)
	session := tts.NewSession(vc, s, engine, outdir)
	session.Open()

	var remove func()
	remove = s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.ChannelID != readChannelID {
			return
		}

		if m.Content == ".stop" {
			session.Close()
			remove()
			vc.Disconnect()
		}

		session.RequestTts(readsense.Format(m.Message), tts.SynthesizeSettings{
			VoiceModel:         "/usr/share/hts-voice/mei_angry.htsvoice",
			Speed:              2,
			AdditionalHalfTone: 0.0,
			AllPassAlpha:       0.5,
			LogF0GVWeight:      1.0,
		})
	})
}
