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
  //endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist?start=" + string(start) + "&end=" + string(end)
	endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist"

	fmt.Println(endpoint)

	data, err := p.Request(NewRequestWithMethod(endpoint, "GET"))

	if err != nil {
		fmt.Println("Endpoint Failed")
		return []*bitfinex.Candle{}, err
	}

	fmt.Println(reflect.TypeOf(data))
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data[0]))
	resolution, err := bitfinex.CandleResolutionFromString(timeframe)
  num := len(data)
  res = make([]*bitfinex.Candle, num)

  for i := 0; i < num; i++ {
			fmt.Println(i)
		  fmt.Println(reflect.TypeOf(data[i]))
			cdl, err := bitfinex.NewCandleFromRaw(symbol, resolution, data[i].(bitfinex.Candle))
			res[i] = cdl
  }

	return res, nil
}
