package main

import (
	"github.com/ardin2001/backend-pemilu/middlewares"
	"github.com/ardin2001/backend-pemilu/routes"
	"github.com/ardin2001/backend-pemilu/utils"
)

func main() {
	utils.MigrateDB()
	e := routes.StartApp()
	middlewares.Logger(e)
	e.Logger.Fatal(e.Start(":8080"))
}
