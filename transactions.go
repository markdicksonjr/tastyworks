package tastyworks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetTransactionsResult struct {
	Data GetTransactionsResultData
	Error ErrorResult
}

type GetTransactionsResultData struct {
	Items []Transaction
}

// BaseTransaction contains fields present on all transaction types
type BaseTransaction struct {
	Id                 int
	AccountNumber      string `json:"account-number"`
	TransactionType    string `json:"transaction-type"`     // e.g. "Money Movement"
	TransactionSubType string `json:"transaction-sub-type"` // e.g. "Balance Adjustment"
	Description        string `json:"description"`          // e.g. "Regulatory fee adjustment"
	ExecutedAt         string `json:"executed-at"`          // e.g. "2021-05-01T13:54:56.792+00:00"
	TransactionDate    string `json:"transaction-date"`     // e.g. "2021-05-01"
	Value              string
	ValueEffect        string `json:"value-effect"` // e.g. "Debit", "Credit"
	NetValue           string `json:"net-value"`
	NetValueEffect     string `json:"net-value-effect"` // e.g. "Debit", "Credit"
	IsEstimatedFee     bool   `json:"is-estimated-fee"`
}

// TradeTransaction reflects the fields only present on a transaction that represents a trade
type TradeTransaction struct {
	Symbol                           string
	InstrumentType                   string `json:"instrument-type"` // e.g. "Equity Option"
	UnderlyingSymbol                 string `json:"underlying-symbol"`
	Action                           string // e.g. "Sell to Open"
	Quantity                         string
	Price                            string
	ClearingFees                     string `json:"clearing-fees"`
	ClearingFeesEffect               string `json:"clearing-fees-effect"` // e.g. "Debit"
	Commission                       string
	CommissionEffect                 string `json:"commission-effect"`
	ProprietaryIndexOptionFees       string `json:"proprietary-index-option-fees"`
	ProprietaryIndexOptionFeesEffect string `json:"proprietary-index-option-fees-effect"`
	ExtExchangeOrderNumber           string `json:"ext-exchange-order-number"`
	ExtGlobalOrderNumber             int    `json:"ext-global-order-number"`
	ExtGroupId                       string `json:"ext-group-id"`
	ExtGroupFillId                   string `json:"ext-group-fill-id"`
	ExtExecId                        string `json:"ext-exec-id"`
	ExecId                           string `json:"exec-id"`
	Exchange                         string // e.g. "1"
	OrderId                          int    `json:"order-id"`
	ExchangeAffiliationIdentifier    string `json:"exchange-affiliation-identifier"`
}

// Transaction contains all fields from BaseTransaction, plus those on specialized transaction types
type Transaction struct {
	BaseTransaction
	TradeTransaction
}

func GetTransactions(sessionToken, accountId string) (GetTransactionsResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/accounts/"+accountId+"/transactions", nil)

	if err != nil {
		return GetTransactionsResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetTransactionsResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetTransactionsResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetTransactionsResult{}, err
	}

	v := GetTransactionsResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetTransactionsResult{}, err
	}

	return v, nil
}
