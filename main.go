package main

import (
	"log"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Info(err)
	}
	return 0;
}
