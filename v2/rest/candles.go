package rest

import (
	"fmt"
	"github.com/dannyluong408/bitfinex-api-go/v2"
	"strings"
	"strconv"
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

	resolution, err := bitfinex.CandleResolutionFromString(timeframe)
	if err != nil {
		fmt.Println("Endpoint Failed2")
		return []*bitfinex.Candle{}, err
	}
  num := len(data)
  res = make([]*bitfinex.Candle, num)

  for i := 0; i < num; i++ {
			fields := strings.Fields(fmt.Sprintf("%v", data[i]))
			for _, res := range fields{
				fmt.Println(res)
			}
			res[i] = &bitfinex.Candle{
				Symbol:     symbol,
				Resolution: resolution,
				MTS:        strconv.ParseInt(fields[2], 64),
				Open:       strconv.ParseFloat(fields[3], 64),
				Close:      strconv.ParseFloat(fields[4], 64),
				High:       strconv.ParseFloat(fields[5], 64),
				Low:        strconv.ParseFloat(fields[6], 64),
				Volume:     strconv.ParseFloat(fields[7], 64),
			}
  }

	return res, nil
}
