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
			fmt.Println(data[i])

			s := make([]string, len(data[i]))
			for i, v := range t {
			    s[i] = fmt.Sprint(v)
			}
			fmt.Println(s)
			// converted := data[i].([]string)[0]
			// fmt.Println(converted)
			// trimmed := converted[1:len(converted)-1]
			// fields := strings.Fields(fmt.Sprintf("%v", trimmed))
			fields := "hello"
			fmt.Println(fields)

			timestamp, err := strconv.ParseInt(fields[0], 10, 64)
			if err != nil {
				fmt.Println("parse int ts failed")
				return []*bitfinex.Candle{}, err
			}
			open, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				fmt.Println("parse float open failed")
				return []*bitfinex.Candle{}, err
			}
			high, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				fmt.Println("parse float high failed")
				return []*bitfinex.Candle{}, err
			}
			low, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				fmt.Println("parse float low failed")
				return []*bitfinex.Candle{}, err
			}
			close, err := strconv.ParseFloat(fields[4], 64)
			if err != nil {
				fmt.Println("parse float close failed")
				return []*bitfinex.Candle{}, err
			}
			vol, err := strconv.ParseFloat(fields[5], 64)
			if err != nil {
					fmt.Println("parse float vol failed")
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
