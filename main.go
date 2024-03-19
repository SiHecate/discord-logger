package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"discord-logger/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutting down or not")
	BotToken       string
)

func main() {
	fmt.Println("Welcome to discord logger bot")

	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()

	BotToken = os.Getenv("BOT_TOKEN")

	if BotToken == "" {
		log.Fatal("Bot access token is required")
	}

	// Discord oturumunu başlat
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %s", err)
	}

	// Olaylarını dinlemek için handler fonksiyonunu çağır
	if err := handler.Handler(dg); err != nil {
		log.Fatalf("Error setting up event handlers: %s", err)
	}

	// Discord sunucusuna katıl
	if err := dg.Open(); err != nil {
		log.Fatalf("Error opening Discord session: %s", err)
	}

	// Botun çalışmasını durdurmak için sinyal işle
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Botu kapat
	if *RemoveCommands {
		if err := dg.Close(); err != nil {
			log.Fatalf("Error closing Discord session: %s", err)
		}
	}
}
