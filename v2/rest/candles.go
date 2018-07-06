package rest

import (
	"github.com/dannyluong408/bitfinex-api-go/utils"
	"fmt"
	"reflect"
)

type CandleService struct {
	Synchronous
}

// Gets KLine History using v2 api
func (p *CandleService) GetOHLCV(timeframe string, symbol string, start int64, end int64) (res []*Candle, err error) {
  endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist?start=" + string(start) + "&end=" + string(end)
	rawdata, err := p.Request(NewRequestWithMethod(endpoint, "GET"))

	if err != nil {
		fmt.Println("Endpoint Failed")
		return []*Candle{}, err
	}

  // fmt.Println("1:")
	// fmt.Println(reflect.TypeOf(data))
	// fmt.Println("2:")
	// fmt.Println(data)
	// fmt.Println("3:")
	// fmt.Println(len(data))
	// fmt.Println("4:")
	// fmt.Println(reflect.TypeOf(data[0]))
	// fmt.Println("5:")
	// fmt.Println(data[0])
  num := len(rawdata)
  res = make([]*Candle, num)
  for i := 0; i < num; i++ {
    res[i] = &Candle{
			Timestamp:                i64ValOrZero(rawdata[0]),
			Open:                     sValOrEmpty(rawdata[1]),
			High:                     sValOrEmpty(rawdata[2]),
			Low:                      sValOrEmpty(rawdata[3]),
			Close:                    sValOrEmpty(rawdata[4]),
			Volume:                   sValOrEmpty(rawdata[5]),
		}
  }
	return res, nil
}

// Kline define kline info
type Candle struct {
	Timestamp                int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
}
