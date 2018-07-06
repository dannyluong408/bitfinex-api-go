package rest

import (
	"fmt"
	"github.com/dannyluong408/bitfinex-api-go/v2"
	"reflect"
)

type CandleService struct {
	Synchronous
}

// Gets KLine History using v2 api
func (p *CandleService) GetOHLCV(timeframe string, symbol string, start int64, end int64) (res []*bitfinex.Candle, err error) {
  endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist?start=" + string(start) + "&end=" + string(end)
	data, err := p.Request(NewRequestWithMethod(endpoint, "GET"))

	if err != nil {
		fmt.Println("Endpoint Failed")
		return []*bitfinex.Candle{}, err
	}
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))
	fmt.Println(reflect.TypeOf(data[0]))
	// resolution, err := bitfinex.CandleResolutionFromString(timeframe)
  // num := len(data)
  // res = make([]*bitfinex.Candle, num)
	//
  for _, candle := range data {
		  fmt.Println(reflect.TypeOf(candle))
	// 		cdl, err := bitfinex.NewCandleFromRaw(symbol, resolution, candle)
	// 		res[i] = cdl
  }

	return res, nil
}
