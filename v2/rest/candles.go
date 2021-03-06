package rest

import (
	"fmt"
	"github.com/dannyluong408/bitfinex-api-go/v2"
	"strconv"
)

type CandleService struct {
	Synchronous
}

// Gets KLine History using v2 api
func (p *CandleService) GetOHLCV(timeframe string, symbol string, start int64, end int64) (res []*bitfinex.Candle, err error) {
  endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist?start=" + strconv.FormatInt(start, 10) + "&end=" + strconv.FormatInt(end, 10)
	//fmt.Println("Endpoint called: %s", endpoint)
	raw, err := p.Request(NewRequestWithMethod(endpoint, "GET"))

	if err != nil {
		fmt.Println("Endpoint Failed")
		return []*bitfinex.Candle{}, err
	}

	resolution, err := bitfinex.CandleResolutionFromString(timeframe)
	if err != nil {
		fmt.Println("Candle Resolution Failed")
		return []*bitfinex.Candle{}, err
	}

  candles := make([]*bitfinex.Candle, 0)

  for _, c := range raw {
			candle, err := bitfinex.NewCandleFromRaw(symbol, resolution, c.([]interface {}))
			if err != nil {
				fmt.Println("NewCandleFromRaw Failed")
				return []*bitfinex.Candle{}, err
			}
			//fmt.Println(candle)
			candles = append(candles, candle)
  }

	return candles, nil
}
