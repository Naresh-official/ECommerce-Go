package main

import (
	"fmt"
	"log"

	"github.com/naresh-official/ecommerce_go/configs"
)

func main() {
	configs.LoadEnv()

	cfg, err := configs.LoadConfig()

	if err != nil {
		log.Fatal("Error in Loading Config ", err)
	}

	fmt.Println("CONFIG", cfg)
}
