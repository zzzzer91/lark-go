package lark_markdown

type parserConfig struct {
	imagePath string
}

func newParserConfig() *parserConfig {
	return &parserConfig{}
}

type Option func(s *parserConfig)

func WithImageHost(imagePath string) Option {
	return func(s *parserConfig) {
		s.imagePath = imagePath
	}
}
