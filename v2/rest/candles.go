package rest

import (
	"github.com/dannyluong408/bitfinex-api-go/v2"
	"net/url"
	"path"
	"strconv"
)

type CandleService struct {
	Synchronous
}

// Gets KLine History using v2 api
func (p *CandleService) GetOHLCV(timeframe string, symbol string, start int64, end int64) (res []*Candle, err error) {
  endpoint := "candles/trade:" + timeframe + ":" + symbol + "/hist?start=" + string(start) + "&end=" + string(end)
	data, err := p.Request(NewRequestWithMethod(endpoint, "GET"))

	if err != nil {
		return []*Candle{}, err
	}
  num := len(data)
  res = make([]*Candle, num)
  for i := 0; i < num; i++ {
    item := data.GetIndex(i)
    res[i] = &Candle{
			Timestamp:                item.GetIndex(0).MustInt64(),
			Open:                     item.GetIndex(1).MustString(),
			High:                     item.GetIndex(2).MustString(),
			Low:                      item.GetIndex(3).MustString(),
			Close:                    item.GetIndex(4).MustString(),
			Volume:                   item.GetIndex(5).MustString(),
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
