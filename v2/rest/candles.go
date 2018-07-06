package rest

import (
	"fmt"
	"github.com/dannyluong408/bitfinex-api-go/v2"
)

type CandleService struct {
	Synchronous
}

// Gets KLine History using v2 api
func (p *CandleService) GetOHLCV(timeframe string, symbol string, start int64, end int64) (res []*bitfinex.Candle, err error) {
  endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist?start=" + string(start) + "&end=" + string(end)
	//endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist"

	fmt.Println(endpoint)

	data, err := p.Request(NewRequestWithMethod(endpoint, "GET"))

	if err != nil {
		fmt.Println("Endpoint Failed")
		return []*bitfinex.Candle{}, err
	}

	resolution, err := bitfinex.CandleResolutionFromString(timeframe)
	if err != nil {
		fmt.Println("Candle Resolution Failed")
		return []*bitfinex.Candle{}, err
	}
	
  res = make([]*bitfinex.Candle, 0)

  for i := 0; i < num; i++ {
			converted := data[i].([]interface{})
			fmt.Println(converted)

			candle, err := bitfinex.NewCandleFromRaw(symbol, resolution, converted)
			if err != nil {
				fmt.Println("NewCandleFromRaw Failed")
				return []*bitfinex.Candle{}, err
			}
			fmt.Println(candle)
			res = append(res, candle)
  }

	return res, nil
}
