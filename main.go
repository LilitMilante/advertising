package main

import (
	"log"

	"github.com/LilitMilante/advertising/internal/app"
)

func main() {
	c := app.NewConfig()

	a, err := app.NewApp(*c)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Start()
	if err != nil {
		log.Fatal(err)
	}
}
