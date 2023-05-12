package main

import (
	"party-calc/readers"
	"party-calc/utils"
)

func main() {
	utils.LoadConfig()
	utils.IntializeLogger()
	readers.StartServer()
}
