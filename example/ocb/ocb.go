package main

import (
	"context"
	"hareta/plugin/ocb"
	"log"
)

func main() {
	bank := ocb.NewOcbBanking("biotductoan", "Toan2004@")
	if err := bank.Login(context.Background()); err != nil {
		log.Fatalln(err)
	}
	if err := bank.List(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
