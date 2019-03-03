package main

import (
	"github.com/andrzejd-pl/SimpleRestBlogBackend/server"
	"github.com/andrzejd-pl/SimpleRestBlogBackend/usage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	usage.CheckErr(err)
	server.Run()
}
