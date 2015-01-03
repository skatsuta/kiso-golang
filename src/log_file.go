package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Print(err)
		return
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("test")

	logger := log.New(file, "[mylogger]", log.LstdFlags)
	logger.Print("test")
}
