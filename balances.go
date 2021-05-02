package tastyworks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetAccountBalanceResult struct {
	Data       GetAccountBalanceResultData
	ApiVersion string `json:"api_version"`
	Context    string `json:"context"`
}

type GetAccountBalanceResultData struct {
	AccountNumber                     string `json:"account-number"`
	CashBalance                       string `json:"cash-balance"`
	LongEquityValue                   string `json:"long-equity-value"`
	ShortEquityValue                  string `json:"short-equity-value"`
	LongDerivativeValue               string `json:"long-derivative-value"`
	ShortDerivativeValue              string `json:"short-derivative-value"`
	LongFuturesValue                  string `json:"long-futures-value"`
	ShortFuturesValue                 string `json:"short-futures-value"`
	LongFuturesDerivativeValue        string `json:"long-futures-derivative-value"`
	ShortFuturesDerivativeValue       string `json:"short-futures-derivative-value"`
	DebitMarginBalance                string `json:"debit-margin-balance"`
	LongMargineableValue              string `json:"long-margineable-value"`
	ShortMargineableValue             string `json:"short-margineable-value"`
	MarginEquity                      string `json:"margin-equity"`
	EquityBuyingPower                 string `json:"equity-buying-power"`
	DerivativeBuyingPower             string `json:"derivative-buying-power"`
	DayTradingBuyingPower             string `json:"day-trading-buying-power"`
	FuturesMarginRequirement          string `json:"futures-margin-requirement"`
	AvailableTradingFunds             string `json:"available-trading-funds"`
	MaintenanceRequirement            string `json:"maintenance-requirement"`
	MaintenanceCallValue              string `json:"maintenance-call-value"`
	RegTCallValue                     string `json:"reg-t-call-value"`
	DayTradingCallValue               string `json:"day-trading-call-value"`
	DayEquityCallValue                string `json:"day-equity-call-value"`
	NetLiquidatingValue               string `json:"net-liquidating-value"`
	CashAvailableToWithdraw           string `json:"cash-available-to-withdraw"`
	DayTradeExcess                    string `json:"day-trade-excess"`
	PendingCash                       string `json:"pending-cash"`
	PendingCashEffect                 string `json:"pending-cash-effect"`
	LongCryptocurrencyValue           string `json:"long-cryptocurrency-value"`
	ShortCryptocurrencyValue          string `json:"short-cryptocurrency-value"`
	CryptocurrencyMarginRequirement   string `json:"cryptocurrency-margin-requirement"`
	UnsettledCryptocurrencyFiatAmount string `json:"unsettled-cryptocurrency-fiat-amount"`
	UnsettledCryptocurrencyFiatEffect string `json:"unsettled-cryptocurrency-fiat-effect"`
	SnapshotDate                      string `json:"snapshot-date"`
	RegTMarginRequirement             string `json:"reg-t-margin-requirement"`
	FuturesOvernightMarginRequirement string `json:"futures-overnight-margin-requirement"`
	FuturesIntradayMarginRequirement  string `json:"futures-intraday-margin-requirement"`
	MaintenanceExcess                 string `json:"maintenance-excess"`
	PendingMarginInterest             string `json:"pending-margin-interest"`
}

func GetAccountBalance(sessionToken, accountId string) (GetAccountBalanceResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/accounts/"+accountId+"/balances", nil)

	if err != nil {
		return GetAccountBalanceResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetAccountBalanceResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetAccountBalanceResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetAccountBalanceResult{}, err
	}

	v := GetAccountBalanceResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetAccountBalanceResult{}, err
	}

	return v, nil
}
