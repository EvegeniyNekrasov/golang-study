package main

import (
	"fmt"

	"github.com/EvegeniyNekrasov/go/cryptomasters/api"
)

func main() {
	ans, _ := api.GetRate("BTC")
	fmt.Println(ans)
}