package openjtalk

type VoiceModel struct {
	filename  string
	modelname string
}

func newVoiceModel(filename string, modelname string) VoiceModel {
	return VoiceModel{
		filename,
		modelname,
	}
}
