package main

import "github.com/chiquitav2/journalful/internal/app"

func main() {
	appSvr := app.NewApp()
	err := appSvr.Init()
	if err != nil {
		panic("Failed to initialize appSvr: " + err.Error())
	}
	appSvr.Start()

}
