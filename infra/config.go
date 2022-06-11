package infra

type FavoriteConfig struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Path   string `yaml:"path"`
}

type YakkyubinConfig struct {
	Scheme  string `yaml:"scheme"`
	Host    string `yaml:"host"`
	Path    string `yaml:"path"`
	ChainId int    `yaml:"chainId"`
}

type Config struct {
	Favorite  FavoriteConfig
	Yakyuubin YakkyubinConfig
}
