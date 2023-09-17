package main

import (
	"fmt"
	"sync"

	"github.com/EvegeniyNekrasov/go/cryptomasters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup
	
	for _, currency := range currencies {
		wg.Add(1)
		go func (currencyCode string) { // immediately invoked function
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	ans, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate fo %v is %.2f\n", ans.Currency, ans.Price)
	}
}