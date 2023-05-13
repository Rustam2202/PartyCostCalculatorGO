package main

import (
	"party-calc/readers"
	"party-calc/utils"
)

func main() {
	utils.IntializeLogger()
	utils.LoadConfig()
	readers.StartServer()
}
