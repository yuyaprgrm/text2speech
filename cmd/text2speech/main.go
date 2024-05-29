package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/yuyaprgrm/text2speech/internal/app/text2speech"
)

func main() {
	primary, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		panic(err)
	}
	// secondary, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN_SUB"))
	// if err != nil {
	// 	panic(err)
	// }

	// pool := dgoworker.NewPool()
	// pool.AddWorker(primary)
	// pool.AddWorker(secondary)
	// s, ok := pool.Checkout()
	// if !ok {
	// 	panic("no worker available")
	// }
	// s.Upgrade().AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
	// 	slog.Info("ready", "user", r.User.ID)
	// })
	// s.Upgrade().Open()
	// defer s.Upgrade().Close()
	// defer s.Return()

	// primary.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
	// 	slog.Info("ready", "user", r.User.ID)
	// })

	// primary.AddHandler(func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	// 	if v.Member.User.Bot {
	// 		return
	// 	}

	// 	if v.ChannelID == "" {
	// 		return
	// 	}

	// 	startTts(s, v.GuildID, v.ChannelID, v.ChannelID)
	// })
	// primary.Open()

	t2s := text2speech.NewText2Speech(primary)
	t2s.Open()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
