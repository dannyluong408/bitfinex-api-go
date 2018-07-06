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
			fields := strings.Fields(fmt.Sprintf("%v", data[i][1:len(data[i])-1]))
			for _, res := range fields{
				fmt.Println(res)
			}
			timestamp, err := strconv.ParseInt(fields[2], 10, 64)
			if err != nil {
				fmt.Println("Endpoint Failed ts")
				return []*bitfinex.Candle{}, err
			}
			open, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				fmt.Println("Endpoint Failed open")
				return []*bitfinex.Candle{}, err
			}
			high, err := strconv.ParseFloat(fields[4], 64)
			if err != nil {
				fmt.Println("Endpoint Failed high")
				return []*bitfinex.Candle{}, err
			}
			low, err := strconv.ParseFloat(fields[5], 64)
			if err != nil {
				fmt.Println("Endpoint Failed low")
				return []*bitfinex.Candle{}, err
			}
			close, err := strconv.ParseFloat(fields[6], 64)
			if err != nil {
				fmt.Println("Endpoint Failed close")
				return []*bitfinex.Candle{}, err
			}
			vol, err := strconv.ParseFloat(fields[7], 64)
			if err != nil {
					fmt.Println("Endpoint Failed vol")
					return []*bitfinex.Candle{}, err
				}

			res[i] = &bitfinex.Candle{
				Symbol:     symbol,
				Resolution: resolution,
				MTS:        timestamp,
				Open:       open,
				Close:      close,
				High:       high,
				Low:        low,
				Volume:     vol,
			}
  }

	return res, nil
}
