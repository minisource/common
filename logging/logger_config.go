package logging

type LoggerConfig struct {
	FilePath string `env:"LOGGER_FILE_PATH"`
	Encoding string `env:"LOGGER_ENCODING"`
	Level    string `env:"LOGGER_LEVEL"`
	Logger   string `env:"LOGGER_LOGGER"`
}
