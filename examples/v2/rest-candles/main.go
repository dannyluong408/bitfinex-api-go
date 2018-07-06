package main

import (
	"flag"
	"fmt"
	"github.com/dannyluong408/bitfinex-api-go/v2/rest"
)

var (
	api     = flag.String("api", "https://api.bitfinex.com/v2/", "v2 REST API URL")
)

// Set BFX_APIKEY and BFX_SECRET as :
//
// export BFX_API_KEY=YOUR_API_KEY
// export BFX_API_SECRET=YOUR_API_SECRET
//
// you can obtain it from https://www.bitfinex.com/api

func main() {
	flag.Parse()

	c := rest.NewClientWithURL(*api)
	fmt.Println("Testing...")
	timeframe := "1m"
	symbol := "tBTCUSD"
	start := int64(516435200000)
	end := int64(1516867200000)

	result, err := c.Candles.GetOHLCV(timeframe, symbol, start, end)
	fmt.Println(result)
	// if err != nil {
	// 	log.Fatalf("getting status: %s", err)
	// }
	//
	// if !available {
	// 	log.Fatalf("API not available")
	// }
	//
	// if *orderid != "" {
	// 	ordid, err := strconv.ParseInt(*orderid, 10, 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	os, err := c.Orders.OrderTrades(bitfinex.TradingPrefix+bitfinex.BTCUSD, ordid)
	// 	if err != nil {
	// 		log.Fatalf("getting order trades: %s", err)
	// 	}
	//
	// 	log.Printf("order trades: %#v\n", os)
	// } else {
	// 	os, err := c.Orders.History(bitfinex.TradingPrefix + bitfinex.BTCUSD)
	// 	if err != nil {
	// 		log.Fatalf("getting orders: %s", err)
	// 	}
	//
	// 	log.Printf("orders: %#v\n", os)
	// }
}
