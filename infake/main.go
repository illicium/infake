package main

import (
	"log"

	"github.com/ypomortsev/infake/infake/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
