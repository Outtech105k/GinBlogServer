package main

import (
	"log"
)

func main() {
	err := serveHTTP()
	if err != nil {
		log.Panic(err)
	}
}
