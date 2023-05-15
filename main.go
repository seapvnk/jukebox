package main

import (
	db "github.com/seapvnk/jukebox/infrastructure/database"
	api "github.com/seapvnk/jukebox/infrastructure/api"
	mig "github.com/seapvnk/jukebox/infrastructure/data"
)

func main() {
	db.Connect()
	mig.Migration()
	api.HandleHttp()
}