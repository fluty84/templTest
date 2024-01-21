package main

import (
	"context"
	"os"
	"templTest/templates"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Hello World")

	helloComponent := templates.Hello("Edu")
	helloComponent.Render(context.Background(), os.Stdout)
}