package config

type Config struct {
	BOT bot `toml:"bot"`
}

type database struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type bot struct {
	BotId  string `toml:"bot_id"`
	ApiKey string `toml:"api_key"`
}
