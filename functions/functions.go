package functions

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	discordBotName = os.Getenv("BOT_NAME")
	logChannelID   = os.Getenv("LOG_CHANNEL_ID")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	discordBotName = os.Getenv("BOT_NAME")
	logChannelID = os.Getenv("LOG_CHANNEL_ID")
	if discordBotName == "" || logChannelID == "" {
		log.Fatal("BOT_NAME and LOG_CHANNEL_ID environment variables must be set")
	}
}

// Mesaj geldiğinde çalışacak fonksiyon
func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Username != discordBotName {
		UserId := m.Author.ID
		ChannelId := m.ChannelID
		SendMessageToLogChannel(s, fmt.Sprintf("Message from:  %s\nUser_ID: %s\nMessage Content: %s\nTime: %s\nChannel_ID: %s\n", m.Author.Mention(), UserId, m.Content, m.Timestamp, ChannelId))
	}
}

// Mesaj silindiğinde çalışacak fonksiyon
func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	if m != nil && m.ID != "" && m.ChannelID != "" {
		SendMessageToLogChannel(s, fmt.Sprintf("Message deleted in channel %s by unknown user", m.ChannelID))
	}
}

// Mesaj güncellediğinde çalışacak fonksiyon
func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.Author.Username != discordBotName {
		UserId := m.Author.ID
		ChannelId := m.ChannelID
		SendMessageToLogChannel(s, fmt.Sprintf("Message updated by:  %s\nUser_ID: %s\nMessage Content: %s\nTime: %s\nChannel_ID: %s\n", m.Author.Mention(), UserId, m.Content, m.Timestamp, ChannelId))
	}
}

// Üye sunucuya katıldığında çalışacak fonksiyon
func OnMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if m.User.Username != discordBotName {
		// Kullanıcının avatar URL'sini oluştur
		SendMessageToLogChannel(s, fmt.Sprintf("%s joined the server", m.User.Username))
	}
}

// Üye sunucudan ayrıldığında çalışacak fonksiyon
func OnMemberLeave(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	if m.User.Username != discordBotName {
		// Kullanıcının avatar URL'sini oluştur
		SendMessageToLogChannel(s, fmt.Sprintf("%s left the server", m.User.Username))
	}
}

// Yeni bir kanal yaratıldığında çalışacak fonksiyon
func OnChannelCreate(s *discordgo.Session, c *discordgo.ChannelCreate) {
	createdBy := c.OwnerID
	channelName := c.Name
	SendMessageToLogChannel(s, fmt.Sprintf("New channel created: %s by %s", channelName, createdBy))
}

// Bir kanal silindiğinde çalışacak fonksiyon
func OnChannelDelete(s *discordgo.Session, c *discordgo.ChannelDelete) {
	channelName := c.Name
	SendMessageToLogChannel(s, fmt.Sprintf("Channel deleted: %s", channelName))
}

// Log kanalına mesaj gönderen yardımcı fonksiyon
func SendMessageToLogChannel(s *discordgo.Session, message string) {
	if logChannelID != "" {
		// Metin kutusunu oluştur
		logMessage := ">>> " + message

		// Kanala mesaj gönder
		_, err := s.ChannelMessageSend(logChannelID, logMessage)
		if err != nil {
			log.Printf("Error sending message to log channel: %s", err)
		}
	}
}
