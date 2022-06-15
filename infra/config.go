package infra

type FavoriteConfig struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Path   string `yaml:"path"`
}

type YakkyubinConfig struct {
	Url string `yaml:"url"`
}

type Config struct {
	Favorite  FavoriteConfig  `yaml:"favorite"`
	Yakyuubin YakkyubinConfig `yaml:"yakkyubin"`
}
