package openjtalk

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	ErrNoVoiceModelsAvailable = errors.New("no voice models available")
)

func VerifyDefaultOpenJtalkInstallation() error {
	ojt, err := exec.LookPath("open_jtalk")
	if err != nil {
		return err
	}

	ojt_proc := exec.Command(ojt)
	if err := ojt_proc.Run(); err != nil {
		return err
	}

	return nil
}

func FindVoiceModels(dir string) ([]VoiceModel, error) {

	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return nil, err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	models := make([]VoiceModel, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".htsvoice") {
			continue
		}

		fullfilename := entry.Name()
		filename := path.Base(fullfilename)
		modelname := strings.TrimSuffix(filename, ".htsvoice")

		models = append(models, newVoiceModel(fullfilename, modelname))
	}

	if len(models) <= 0 {
		return nil, ErrNoVoiceModelsAvailable
	}

	return models, nil
}

func VerifyDefaultVoiceModels() error {

	models, err := FindVoiceModels("/usr/share/hts-voice")

	if err != nil {
		return err
	}

	log.Println("found voice models", "models", models)

	return nil
}
