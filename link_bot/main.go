package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	discordinterface "link_bot/discordInterface"

	"github.com/spf13/viper"
)

type Config struct {
	DiscordToken string
}

func loadConfig() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, fmt.Errorf("read .env: %w", err)
		}
	}

	config := Config{
		DiscordToken: viper.GetString("DISCORD_TOKEN"),
	}
	if config.DiscordToken == "" {
		return Config{}, fmt.Errorf("DISCORD_TOKEN is required")
	}
	return config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	bot, err := discordinterface.NewBot(discordinterface.BotConfig{
		Token: config.DiscordToken,
	})
	if err != nil {
		slog.Error("failed to create Discord bot", "error", err)
		os.Exit(1)
	}
	if err := bot.Run(); err != nil {
		slog.Error("Discord bot stopped", "error", err)
		os.Exit(1)
	}
}
