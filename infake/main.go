package main

import (
	"log"

	"revision.aeip.apigee.net/dia/infake/infake/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
