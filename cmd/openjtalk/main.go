package main

import (
	"bytes"
	"log/slog"

	"github.com/yuyaprgrm/text2speech/internal/pkg/openjtalk"
)

func main() {
	models, err := openjtalk.FindVoiceModels("/usr/share/hts-voice")
	if err != nil {
		panic(err)
	}

	slog.Info("found voice models", "models", models)

	ctx := openjtalk.NewCtx().WithDictionaryDirectory("/var/lib/mecab/dic/open-jtalk/naist-jdic")

	cmd := openjtalk.BuildCommand("/usr/bin/open_jtalk", "こんにちは", ctx)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		slog.Error("failed to synthesize", "error", stderr.String())
		panic(err)
	}
}
