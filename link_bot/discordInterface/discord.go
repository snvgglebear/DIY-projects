package discordinterface

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"github.com/bwmarrin/discordgo"
)

type BotConfig struct {
	Token string
}

type Bot struct {
	session *discordgo.Session
}

func NewBot(config BotConfig) (*Bot, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, fmt.Errorf("create Discord session: %w", err)
	}
	session.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		slog.Info("Discord bot connected", "user", ready.User.Username)
	})
	return &Bot{session: session}, nil
}

func (bot *Bot) Run() error {
	if err := bot.session.Open(); err != nil {
		return fmt.Errorf("open Discord session: %w", err)
	}
	defer bot.session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	return nil
}
