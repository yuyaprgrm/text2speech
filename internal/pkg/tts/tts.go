package tts

import (
	"log/slog"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

type SynthesizeSettings struct {
	VoiceModel         string
	Speed              float64
	AdditionalHalfTone float64
	AllPassAlpha       float64
	LogF0GVWeight      float64
}

type synthesizeRequest struct {
	text     string
	settings SynthesizeSettings
}

type playRequest struct {
	file string
}

type Session struct {
	vc              *discordgo.VoiceConnection
	dg              *discordgo.Session
	engine          *Engine
	outdir          string
	synthesizeQueue chan synthesizeRequest
	playQueue       chan playRequest
	close           chan struct{}
}

func NewSession(vc *discordgo.VoiceConnection, dg *discordgo.Session, engine *Engine, outdir string) *Session {
	return &Session{
		vc:              vc,
		dg:              dg,
		engine:          engine,
		outdir:          outdir,
		synthesizeQueue: make(chan synthesizeRequest, 8),
		playQueue:       make(chan playRequest, 8),
	}
}

func (s *Session) RequestTts(text string, settings SynthesizeSettings) {
	s.synthesizeQueue <- synthesizeRequest{
		text,
		settings,
	}
}

func (s *Session) Open() {

	// synthesize
	go func() {
		for {
			select {
			case <-s.close:
				return
			case request := <-s.synthesizeQueue:

				id := uuid.New()
				target := s.outdir + "/" + id.String() + ".wav"
				sanitized := request.text
				err := s.engine.Synthesize(sanitized, target, &request.settings)
				if err != nil {
					slog.Debug("failed to generate", err)
					continue
				}
				s.playQueue <- playRequest{target}
			}
		}
	}()

	// play
	go func() {
		for {
			select {
			case <-s.close:
				return
			case request := <-s.playQueue:
				dgvoice.PlayAudioFile(s.vc, request.file, make(chan bool))
				os.Remove(request.file)
			}
		}
	}()
}

func (s *Session) Close() {
	s.close <- struct{}{}
	close(s.close)
}
