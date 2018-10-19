package main

import (
	"log"
	"os"
)

func main() {
	env_username := os.Getenv("username")
	log.Println("env_username=", env_username)
}
