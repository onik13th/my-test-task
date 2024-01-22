package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

// структура Config для настройки сервера API.
// поля имеют свои аннотации для декодирования значений из файла конфигурации с использованием пакета toml.

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

// NewConfig() возвращает указатель на новый объект Config
// со значениями по умолчанию для BindAddr и LogLevel.
