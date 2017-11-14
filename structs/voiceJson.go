package structs

//语音识别 request data
type VoiceJson struct {
	Format string `json:"format"`
	Rate int `json:"rate"`
	Channel int `json:"channel"`
	Cuid string `json:"cuid"`
	Token string `json:"token"`
	Speech string `json:"speech"`
	Len int `json:"len"`
}
