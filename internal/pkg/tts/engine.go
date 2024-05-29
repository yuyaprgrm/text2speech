package tts

import (
	"bytes"
	"errors"

	"github.com/yuyaprgrm/text2speech/pkg/openjtalk"
)

// Engine is the struct that holds the configuration for the OpenJTalk engine.
type Engine struct {
	enginePath               string
	defaultSynthesizeContext *openjtalk.SynthesizeCtx
}

func NewEngine(enginePath string, defaultSynthesizeContext *openjtalk.SynthesizeCtx) *Engine {
	return &Engine{
		enginePath,
		defaultSynthesizeContext,
	}
}

func (engine *Engine) Synthesize(text string, destination string, settings *SynthesizeSettings) error {
	ctx := engine.defaultSynthesizeContext.WithWavOutput(destination)

	// voice settings
	ctx = ctx.WithVoiceModel(settings.VoiceModel).
		WithSpeechSpeed(settings.Speed).
		WithAdditionalHalfTone(settings.AdditionalHalfTone).
		WithAllPassAlpha(settings.AllPassAlpha).
		WithLogF0GVWeight(settings.LogF0GVWeight)

	run := openjtalk.BuildCommand(engine.enginePath, text, ctx)

	var stderr bytes.Buffer
	run.Stderr = &stderr
	run.Start()

	if run.Wait() != nil {
		return errors.New("synthesize error: " + stderr.String())
	}
	return nil
}

func (engine *Engine) AsyncSynthesize(text string, destination string, settings *SynthesizeSettings) <-chan error {
	ch := make(chan error)
	go func() {
		ch <- engine.Synthesize(text, destination, settings)
		close(ch)
	}()
	return ch
}
