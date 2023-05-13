package main

import (
	"party-calc/internal/web"
	"party-calc/utils"
)

func main() {
	utils.IntializeLogger()
	utils.LoadConfig()
	web.StartServer()
}
