package infra

type ApiConfig struct {
	Url string `yaml:"url"`
}

type Config struct {
	Favorite  ApiConfig `yaml:"favorite"`
	Yakyuubin ApiConfig `yaml:"yakkyubin"`
}
