package main

import (
	"github.com/joho/godotenv"
	"mysql/server"
	"mysql/usage"
)

func main() {
	err := godotenv.Load(".env")
	usage.CheckErr(err)
	server.Run()
}
