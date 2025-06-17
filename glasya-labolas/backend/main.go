package main

import (
	"glasya-labolas/cmd"
	"glasya-labolas/database"
)

func main() {
	database.InitDB()
	cmd.Execute()
}