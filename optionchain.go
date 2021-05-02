package tastyworks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetOptionChainForTickerResult struct {
	Data  GetOptionChainForTickerResultData
	Error ErrorResult
}

type GetOptionChainForTickerResultData struct {
	Items []OptionChainItem
}

type OptionChainItem struct {
	Symbol                         string
	InstrumentType                 string `json:"instrument-type"`
	Active                         bool
	StrikePrice                    string `json:"strike-price"`
	RootSymbol                     string `json:"root-symbol"`
	UnderlyingSymbol               string `json:"underlying-symbol"`
	ExpirationDate                 string `json:"expiration-date"`
	ExerciseStyle                  string `json:"exercise-style"`
	SharesPerContract              int    `json:"shares-per-contract"`
	OptionType                     string `json:"option-type"`
	OptionChainType                string `json:"option-chain-type"`
	ExpirationType                 string `json:"expiration-type"`                   // e.g. Regular
	SettlementType                 string `json:"settlement-type"`                   // e.g. PM
	StopsTradingAt                 string `json:"stops-trading-at"`                  // e.g. "2022-02-18T21:15:00.000+00:00"
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"` // e.g. "Cash Settled Equity Option"
	DaysToExpiration               int    `json:"days-to-expiration"`
	ExpiresAt                      string `json:"expires-at"`
	IsClosingOnly                  bool   `json:"is-closing-only"`
}

type GetNestedOptionChainForTickerResult struct {
	Data    GetNestedOptionChainForTickerResultData
	Error   ErrorResult
	Context string
}

type GetNestedOptionChainForTickerResultData struct {
	Items []NestedOptionChainItem
}

type NestedOptionChainItem struct {
	UnderlyingSymbol  string     `json:"underlying-symbol"`
	RootSymbol        string     `json:"root-symbol"`
	OptionChainType   string     `json:"option-chain-type"`
	SharesPerContract float64    `json:"shares-per-contract"`
	TickSizes         []TickSize `json:"tick-sizes"`
	Deliverables      []Deliverable
	Expirations       []Expiration
}

type TickSize struct {
	Value string
}

type Deliverable struct {
	Id              int
	RootSymbol      string `json:"root-symbol"`
	DeliverableType string `json:"deliverable-type"` // e.g. "Shares"
	Description     string
	Amount          string
	Symbol          string
	InstrumentType  string `json:"instrument-type"` // e.g. "Equity"
	Percent         string
}

type Expiration struct {
	ExpirationType   string `json:"expiration-type"` // e.g. Weekly
	ExpirationDate   string `json:"expiration-date"` // e.g. "2021-05-03"
	DaysToExpiration int    `json:"days-to-expiration"`
	SettlementType   string `json:"settlement-type"`
	Strikes          []Strike
}

type Strike struct {
	StrikePrice string `json:"strike-price"`
	Call        string // e.g. "SPY   210503C00240000"
	Put         string // e.g. "SPY   210503P00240000"
}

func GetOptionChainForTicker(sessionToken string, ticker string) (GetOptionChainForTickerResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/option-chains/"+ticker, nil)

	if err != nil {
		return GetOptionChainForTickerResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetOptionChainForTickerResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetOptionChainForTickerResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetOptionChainForTickerResult{}, err
	}

	v := GetOptionChainForTickerResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetOptionChainForTickerResult{}, err
	}

	if v.Error.Code != "" {
		return v, errors.New(v.Error.Message)
	}

	return v, nil
}

func GetNestedOptionChainForTicker(sessionToken string, ticker string) (GetNestedOptionChainForTickerResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/option-chains/"+ticker+"/nested", nil)

	if err != nil {
		return GetNestedOptionChainForTickerResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetNestedOptionChainForTickerResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetNestedOptionChainForTickerResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetNestedOptionChainForTickerResult{}, err
	}

	v := GetNestedOptionChainForTickerResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetNestedOptionChainForTickerResult{}, err
	}

	if v.Error.Code != "" {
		return v, errors.New(v.Error.Message)
	}

	return v, nil
}
