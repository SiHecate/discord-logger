package handler

import (
	"discord-logger/functions"

	"github.com/bwmarrin/discordgo"
)

// Handler fonksiyonu, Discord oturumuna event handler'ları eklemek için kullanılır. Bu event handler'lar, çeşitli Discord olaylarını dinlemek ve işlemek için kullanılır.
func Handler(dg *discordgo.Session) error {
	dg.AddHandler(functions.OnMessage)
	dg.AddHandler(functions.OnMemberJoin)
	dg.AddHandler(functions.OnMemberLeave)
	dg.AddHandler(functions.MessageDelete)
	dg.AddHandler(functions.MessageUpdate)
	dg.AddHandler(functions.OnChannelCreate)
	dg.AddHandler(functions.OnChannelDelete)
	return nil
}
