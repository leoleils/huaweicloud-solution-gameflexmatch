package config

var GlobalConfig = &Config{}

type Config struct {
	LogPath                 string
	HttpStartPort           int
	HttpEndPort             int
	GameSessionRetainMinute int
	ProcessRunMinute        int
	MaxGameSessionCount		int

	Ak						string
	Sk						string
	TopicUrn				string
}