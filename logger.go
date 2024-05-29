package main

import (
	"log"
	"os"
  "strings"
)

type LogLevel int

const (
  ERROR LogLevel = iota
  WARNING LogLevel = iota
  INFO LogLevel = iota
  DEBUG LogLevel = iota
)

type Logger struct {
  LOG_LEVEL LogLevel
  i *log.Logger
  e *log.Logger
  d *log.Logger
}

func newLogger() Logger {
  return Logger{
    getLogLevel(),
    log.New(os.Stdout, "INFO -", log.Ldate|log.Ltime|log.Lshortfile),
    log.New(os.Stderr, "ERROR -", log.Ldate|log.Ltime|log.Lshortfile),
    log.New(os.Stdout, "DEBUG -", log.Ldate|log.Ltime|log.Lshortfile),
  }
}

func getLogLevel() LogLevel {
  levelString := os.Getenv("LOG_LEVEL")

  switch strings.ToLower(levelString) {
    case "error":
      return ERROR
    case "warning":
      return WARNING
    case "info":
      return INFO
    case "debug":
      return DEBUG
  }

  return INFO
}

func (l Logger) Error(f string, s ...any) {
  l.i.Printf(f+"\n", s...)
}

func (l Logger) Info(f string, s ...any) {
  if l.LOG_LEVEL >= INFO {
    l.i.Printf(f+"\n", s...)
  }
}

func (l Logger) Debug(f string, s ...any) {
  if l.LOG_LEVEL >= DEBUG {
    l.d.Printf(f+"\n", s...)
  }
}
