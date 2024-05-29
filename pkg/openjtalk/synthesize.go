package openjtalk

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type SynthesizeCtx struct {
	context.Context
}

func NewCtx() *SynthesizeCtx {
	return &SynthesizeCtx{context.Background()}
}

func ctxWith(ctx context.Context) *SynthesizeCtx {
	return &SynthesizeCtx{ctx}
}

func (ctx *SynthesizeCtx) WithDictionaryDirectory(dir string) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "DictionaryDirectory", dir))
}

func (ctx *SynthesizeCtx) WithVoiceModel(model string) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "VoiceModel", model))
}

func (ctx *SynthesizeCtx) WithWavOutput(output string) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "WavOutput", output))
}

func (ctx *SynthesizeCtx) WithTraceOutput(output string) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "TraceOutput", output))
}

func (ctx *SynthesizeCtx) WithSamplingFrequency(frequency float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "SamplingFrequency", frequency))
}

func (ctx *SynthesizeCtx) WithFramePeriod(period float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "FramePeriod", period))
}

func (ctx *SynthesizeCtx) WithAllPassAlpha(alpha float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "AllPassAlpha", alpha))
}

func (ctx *SynthesizeCtx) WithPostfilteringCoeff(coeff float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "PostfilteringCoeff", coeff))
}

func (ctx *SynthesizeCtx) WithSpeechSpeed(speed float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "SpeechSpeed", speed))
}

func (ctx *SynthesizeCtx) WithAdditionalHalfTone(halfTone float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "AdditionalHalfTone", halfTone))
}

func (ctx *SynthesizeCtx) WithVUVThreshold(threshold float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "VUVThreshold", threshold))
}

func (ctx *SynthesizeCtx) WithSpectrumGVWeight(weight float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "SpectrumGVWeight", weight))
}

func (ctx *SynthesizeCtx) WithLogF0GVWeight(weight float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "LogF0GVWeight", weight))
}

func (ctx *SynthesizeCtx) WithVolume(volume float64) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "Volume", volume))
}

func (ctx *SynthesizeCtx) WithAudioBufferSize(bufferSize int) *SynthesizeCtx {
	return ctxWith(context.WithValue(ctx, "AudioBufferSize", bufferSize))
}

func (ctx *SynthesizeCtx) build() (args []string) {
	args = make([]string, 0)
	if x, ok := ctx.Value("DictionaryDirectory").(string); ok {
		args = append(args, "-x", x)
	}

	if m, ok := ctx.Value("VoiceModel").(string); ok {
		args = append(args, "-m", m)
	}

	if w, ok := ctx.Value("WavOutput").(string); ok {
		args = append(args, "-ow", w)
	}
	if t, ok := ctx.Value("TraceOutput").(string); ok {
		args = append(args, "-ot", t)
	}
	if f, ok := ctx.Value("SamplingFrequency").(float64); ok {
		args = append(args, "-s", fmt.Sprintf("%.2f", f))
	}
	if p, ok := ctx.Value("FramePeriod").(float64); ok {
		args = append(args, "-p", fmt.Sprintf("%.2f", p))
	}
	if a, ok := ctx.Value("AllPassAlpha").(float64); ok {
		args = append(args, "-a", fmt.Sprintf("%.2f", a))
	}
	if c, ok := ctx.Value("PostfilteringCoeff").(float64); ok {
		args = append(args, "-b", fmt.Sprintf("%.2f", c))
	}
	if s, ok := ctx.Value("SpeechSpeed").(float64); ok {
		args = append(args, "-r", fmt.Sprintf("%.2f", s))
	}
	if h, ok := ctx.Value("AdditionalHalfTone").(float64); ok {
		args = append(args, "-fm", fmt.Sprintf("%.2f", h))
	}
	if v, ok := ctx.Value("VUVThreshold").(float64); ok {
		args = append(args, "-u", fmt.Sprintf("%.2f", v))
	}
	if g, ok := ctx.Value("SpectrumGVWeight").(float64); ok {
		args = append(args, "-jm", fmt.Sprintf("%.2f", g))
	}
	if l, ok := ctx.Value("LogF0GVWeight").(float64); ok {
		args = append(args, "-jf", fmt.Sprintf("%.2f", l))
	}
	if v, ok := ctx.Value("Volume").(float64); ok {
		args = append(args, "-g", fmt.Sprintf("%.2f", v))
	}
	if b, ok := ctx.Value("AudioBufferSize").(int); ok {
		args = append(args, "-b", fmt.Sprintf("%d", b))
	}

	return args
}

func BuildCommand(ojt_path string, text string, ctx *SynthesizeCtx) *exec.Cmd {
	args := ctx.build()
	cmd := exec.Command(
		ojt_path,
		args...,
	)
	cmd.Stdin = strings.NewReader(text)
	return cmd
}
