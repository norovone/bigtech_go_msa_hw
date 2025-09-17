package main

import (
	"log"

	"github.com/norovone/bigtech_go_msa_hw/gateway/internal/gateway"
)

func main() {
	if err := gateway.Run(); err != nil {
		log.Fatal(err)
	}
}
