package conf

type Config struct {
	Zap ZapConfig
}

type ZapConfig struct {
	Director string
}
