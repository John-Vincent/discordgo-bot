package main

import (
	"os"
  "os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

var (
  logger = newLogger()
)

func main() {
  apiKey := os.Getenv("DISCORD_API_KEY")
  logger.Debug(apiKey)

  dg, err := discordgo.New(apiKey)
  if err != nil {
    logger.Error("error creating discordgo session: %+v", err)
    return
  }

  //attach handlers here
  dg.AddHandler(loginHandler)
  dg.AddHandler(messageEvent)
  dg.AddHandler(pingPong)

  if err := dg.Open(); err != nil {
    logger.Error("failed to open connection to discord: %+v", err)
    return
  }
  defer dg.Close()

  logger.Info("bot has started up")

  //blocks until one of the signals is received
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGSEGV, syscall.SIGHUP)
  <-sc
}

