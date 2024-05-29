package openjtalk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuyaprgrm/text2speech/pkg/openjtalk"
)

func TestSynthesize(t *testing.T) {
	// open_jtalk -m /usr/share/hts-voice/mei_normal.htsvoice -x /var/lib/mecab/dic/open-jtalk/naist-jdic -ow /tmp/out.wav
	ctx := openjtalk.NewCtx().
		WithVoiceModel("/usr/share/hts-voice/mei_normal.htsvoice").
		WithDictionaryDirectory("/var/lib/mecab/dic/open-jtalk/naist-jdic").
		WithWavOutput("/tmp/out.wav")

	cmd := openjtalk.BuildCommand("/usr/bin/open_jtalk", "こんにちは", ctx)
	err := cmd.Run()

	assert.Nil(t, err)
}
