package openjtalk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuyaprgrm/text2speech/pkg/openjtalk"
)

func TestVerifyDefaultOpenJtalkInstallation(t *testing.T) {
	err := openjtalk.VerifyDefaultOpenJtalkInstallation()
	assert.Nil(t, err)
}

func TestVerifyDefaultVoiceModels(t *testing.T) {
	err := openjtalk.VerifyDefaultVoiceModels()
	assert.NotNil(t, err)
}
