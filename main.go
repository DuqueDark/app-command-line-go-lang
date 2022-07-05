package main

import (
	"log"
	"os"

	"github.com/duquedark/app_command_line/app"
)

func main() {
	//fmt.Println("Inicio app")
	app := app.Gerar()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
