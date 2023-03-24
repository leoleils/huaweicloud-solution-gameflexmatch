package config


var GlobalConfig Config

type Config struct {
	LogPath		string
	HttpStartPort	int
	HttpEndPort		int
}