package main

import (
	apiserver "github.com/Andronovdima/AvitoChatAssignment/internal/app"
	"log"
)

var newline string

func main() {
	if err := apiserver.Stop(); err != nil {
		log.Info(err)
	}
	return 0;
}
