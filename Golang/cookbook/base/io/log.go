package io

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	apexlog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/sirupsen/logrus"
)

// Log uses the setup logger
func Log() {
	// we'll configure the logger to write
	// to a bytes.Buffer
	buf := bytes.Buffer{}

	// second argument is the prefix last argument is about options
	// you combine them with a logical or.
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)

	fmt.Println(buf.String())
}

// Hook will implement the logrus
// hook interface
type Hook struct {
	id string
}

// Fire will trigger whenever you log
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

// Levels is what levels this hook will fire on
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus demonstrates some basic logrus functionality
func Logrus() {
	// we're emitting in json format
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
}

// ThrowError throws an error that we'll trace
func ThrowError() error {
	err := errors.New("a crazy failure")
	apexlog.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

// CustomHandler splits to two streams
type CustomHandler struct {
	id      string
	handler apexlog.Handler
}

// HandleLog adds a hook and does the emitting
func (h *CustomHandler) HandleLog(e *apexlog.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

// Apex has a number of useful tricks
func Apex() {
	apexlog.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	//With error convenience function
	apexlog.WithError(err).Error("an error occurred")
}
