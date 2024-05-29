package main

import (
  "strings"
	"github.com/bwmarrin/discordgo"
)

func messageEvent(session *discordgo.Session, message *discordgo.MessageCreate) {
  logger.Debug("message: \"%+v\"", *message.Message)
}

func pingPong(session *discordgo.Session, message *discordgo.MessageCreate) {
  content := strings.TrimSpace(strings.ToLower(message.Content))
  logger.Debug("processing ping: %s", content)

  if message.Author.ID == session.State.User.ID {
    logger.Debug("I sent this message")
    return
  }

  if len(message.GuildID) > 0 {
    logger.Debug("message isn't a dm")
    return
  }

  if content != "!ping" {
    logger.Debug("\"%s\" is not equal to \"%s\"", content, "!ping")
    return
  }

  sentMessage, err := session.ChannelMessageSend(message.ChannelID, "pong")

  if err != nil {
    logger.Error("failed to send message: %+v", err)
    return
  }

  logger.Info("pong sent: %s", sentMessage.ID)
}

func loginHandler(session *discordgo.Session, r *discordgo.Ready) {
  logger.Info("Logged in as: %v#%v", session.State.User.Username, session.State.User.Discriminator)
}
