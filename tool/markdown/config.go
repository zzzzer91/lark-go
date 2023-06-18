package markdown

type parserConfig struct {
	imagePath    string
	isEnableHtml bool
}

func newParserConfig() *parserConfig {
	return &parserConfig{}
}

type Option func(cfg *parserConfig)

func WithImagePath(imagePath string) Option {
	return func(cfg *parserConfig) {
		cfg.imagePath = imagePath
	}
}

func WithIsEnableHtml(flag bool) Option {
	return func(cfg *parserConfig) {
		cfg.isEnableHtml = flag
	}
}
