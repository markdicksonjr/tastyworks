package tastyworks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetPositionsResult struct {
	Data GetPositionsResultData
	Error ErrorResult
}

type GetPositionsResultData struct {
	Items []Position
}

type Position struct {
	AccountNumber                 string `json:"account-number"`
	Symbol                        string
	InstrumentType                string `json:"instrument-type"` // e.g. "Equity Option"
	UnderlyingSymbol              string `json:"underlying-symbol"`
	Quantity                      float64
	QuantityDirection             string `json:"quantity-direction"` // e.g. "Short"
	ClosePrice                    string `json:"close-price"`
	AverageOpenPrice              string `json:"average-open-price"`
	AverageYearlyMarketClosePrice string `json:"average-yearly-market-close-price"`
	AverageDailyMarketClosePrice  string `json:"average-daily-market-close-price"`
	Multiplier                    float64
	CostEffect                    string  `json:"cost-effect"` // e.g. "Debit"
	IsSuppressed                  bool    `json:"is-suppressed"`
	IsFrozen                      bool    `json:"is-frozen"`
	RestrictedQuantity            float64 `json:"restricted-quantity"`
	ExpiresAt                     string  `json:"expires-at"`
	RealizedDayGain               string  `json:"realized-day-gain"`
	RealizedDayGainEffect         string  `json:"realized-day-gain-effect"` // e.g. "None"
	RealizedDayGainDate           string  `json:"realized-day-gain-date"`
	RealizedToday                 string  `json:"realized-today"`
	RealizedTodayEffect           string  `json:"realized-today-effect"` // e.g. "None"
	RealizedTodayDate             string  `json:"realized-today-date"`
	CreatedAt                     string  `json:"created-at"`
	UpdatedAt                     string  `json:"updated-at"`
}

func GetPositions(sessionToken, accountId string) (GetPositionsResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/accounts/"+accountId+"/positions", nil)

	if err != nil {
		return GetPositionsResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetPositionsResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetPositionsResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetPositionsResult{}, err
	}

	v := GetPositionsResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetPositionsResult{}, err
	}

	if v.Error.Code != "" {
		return v, errors.New(v.Error.Message)
	}

	return v, nil
}
