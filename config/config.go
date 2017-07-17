package config

type Config struct {
	BOT bot      `toml:"bot"`
	DB  database `toml:"database"`
}

type database struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Datapath string `toml:"datapath"`
}

type bot struct {
	BotId      string `toml:"bot_id"`
	ApiKey     string `toml:"api_key"`
	MainChatId int    `toml:"main_chat_id"`
}
