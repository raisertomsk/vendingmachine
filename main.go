package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags)

	configPath := flag.String("c", "./config.json", "Configuration file")
	flag.Parse()

	config, err := NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	app := NewApp(config)
	app.Run()
}
