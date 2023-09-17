package main

import (
	"fmt"

	"github.com/EvegeniyNekrasov/go/cryptomasters/api"
)

func main() {
	ans, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("The rate fo %v is %.2f", ans.Currency, ans.Price)
	}
}