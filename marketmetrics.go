package tastyworks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetMarketMetricsResult struct {
	Data  GetMarketMetricsResultData
	Error ErrorResult
}
type GetMarketMetricsResultData struct {
	Items []GetMarketMetricsSymbolResult
}

type GetMarketMetricsSymbolResult struct {
	Symbol                                 string
	ImpliedVolatilityIndex                 string        `json:"implied-volatility-index"`
	ImpliedVolatilityIndex5DayChange       string        `json:"implied-volatility-index-5-day-change"`
	ImpliedVolatilityIndexRank             string        `json:"implied-volatility-index-rank"`
	TOSImpliedVolatilityIndexRank          string        `json:"tos-implied-volatility-index-rank"`
	TWImpliedVolatilityIndexRank           string        `json:"tw-implied-volatility-index-rank"`
	TOSImpliedVolatilityIndexRankUpdatedAt string        `json:"tos-implied-volatility-index-rank-updated-at"`
	TWImpliedVolatilityIndexRankUpdatedAt  string        `json:"tw-implied-volatility-index-rank-updated-at"`
	ImpliedVolatilityIndexRankSource       string        `json:"implied-volatility-index-rank-source"`
	ImpliedVolatilityPercentile            string        `json:"implied-volatility-percentile"`
	ImpliedVolatilityUpdatedAt             string        `json:"implied-volatility-updated-at"`
	LiquidityValue                         string        `json:"liquidity-value"`
	LiquidityRank                          string        `json:"liquidity-rank"`
	LiquidityRating                        int           `json:"liquidity-rating"`
	UpdatedAt                              string        `json:"updated-at"`
	OptionExpirationImpliedVolatilities    []interface{} `json:"option-expiration-implied-volatilities"` // TODO
	LiquidityRunningState                  LiquidityRunningState
	Beta                                   string
	CorrSpy3Month                          string                   `json:"corr-spy-3month"`
	DividendRatePerShare                   string                   `json:"dividend-rate-per-share"`
	AnnualDividendPerShare                 string                   `json:"annual-dividend-per-share"`
	DividendYield                          string                   `json:"dividend-yield"`
	DividendExDate                         string                   `json:"dividend-ex-date"`
	DividendNextDate                       string                   `json:"dividend-next-date"`
	DividendPayDate                        string                   `json:"dividend-pay-date"`   // e.g. "2021-04-30",
	DividendUpdatedAt                      string                   `json:"dividend-updated-at"` // e.g. "2021-03-20T03:15:03.876Z"
	Earnings                               GetMarketMetricsEarnings `json:"earnings"`
	ListedMarket                           string                   `json:"listed-market"`
	Lendability                            string                   `json:"lendability"` // e.g. "Easy To Borrow"
	BorrowRate                             string                   `json:"borrow-rate"` // e.g. "1.25"
}

type GetMarketMetricsEarnings struct {
	Visible            bool
	ExpectedReportDate string `json:"expected-report-date"` // e.g. "2021-04-29"
	Estimated          bool
	TimeOfDay          string `json:"time-of-day"` // e.g. "AMC"
	LateFlag           bool   `json:"late-flag"`
	QuarterEndDate     string `json:"quarter-end-date"`
	ActualEps          string `json:"actual-eps"`
	ConsensusEstimate  string `json:"consensus-estimate"`
	UpdatedAt          string `json:"updated-at"` // e.g. "2021-04-29T11:00:06.668Z"
}

type LiquidityRunningState struct {
	Sum       string
	Count     int
	StartedAt string `json:"started-at"` // e.g. "2021-05-01T10:00:14.756Z"
}

type OptionExpirationImpliedVolatility struct {
	ExpirationDate    string `json:"expiration-date"`
	OptionChainType   string `json:"option-chain-type"`
	SettlementType    string `json:"settlement-type"`
	ImpliedVolatility string `json:"implied-volatility"`
}

func GetMarketMetricsForTickers(sessionToken string, tickers []string) (GetMarketMetricsResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/market-metrics?symbols="+strings.Join(tickers, ","), nil)

	if err != nil {
		return GetMarketMetricsResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetMarketMetricsResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetMarketMetricsResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetMarketMetricsResult{}, err
	}

	v := GetMarketMetricsResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetMarketMetricsResult{}, err
	}

	return v, nil
}
