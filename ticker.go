package hitbtc

import (
	"encoding/json"
	"time"
)

type Ticker struct {
	Ask         string    `json:"ask"`
	Bid         string    `json:"bid"`
	Last        string    `json:"last"`
	Open        string    `json:"open"`
	Low         string    `json:"low"`
	High        string    `json:"high"`
	Volume      string    `json:"volume"`
	VolumeQuote string    `json:"volumeQuote"`
	Timestamp   time.Time `json:"timestamp"`
	Symbol      string    `json:"symbol"`
}

func (t *Ticker) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Ticker
	aux := &struct {
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Timestamp, err = time.Parse("2006-01-02T15:04:05.999Z", aux.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
