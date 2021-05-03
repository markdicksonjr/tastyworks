package tastyworks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetLiveOrdersResult struct {
	Data       GetLiveOrdersResultData
	Error      ErrorResult
	ApiVersion string `json:"api-version"`
	Context    string `json:"context"`
}

type GetLiveOrdersResultData struct {
	Items []Order
}

type Order struct {
	Id                       int
	AccountNumber            string `json:"account-number"`
	TimeInForce              string `json:"time-in-force"` // e.g "GTC"
	OrderType                string `json:"order-type"`
	Size                     int
	UnderlyingSymbol         string `json:"underlying-symbol"`
	UnderlyingInstrumentType string `json:"underlying-instrument-type"`
	Price                    string
	PriceEffect              string `json:"price-effect"` // e.g. "Debit"
	Status                   string // e.g. "Received"
	Cancellable              bool
	Editable                 bool
	Edited                   bool
	ExtExchangeOrderNumber   string `json:"ext-exchange-order-number"`
	ExtClientOrderId         string `json:"ext-client-order-id"`
	ExtGlobalOrderNumber     int    `json:"ext-global-order-number"`
	ReceivedAt               string `json:"received-at"` // e.g. "2021-04-29T18:38:32.506+00:00"
	UpdatedAt                int64  `json:"updated-at"`  // e.g. 1619813203353
	Legs                     []OrderLeg
}

type OrderLeg struct {
	InstrumentType    string
	Symbol            string
	Quantity          float64
	RemainingQuantity float64
	Action            string // e.g. "Buy to Close"
	Fills             []interface{} // TODO: improve type
}

func GetLiveOrders(sessionToken, accountId string) (GetLiveOrdersResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/accounts/"+accountId+"/orders/live", nil)

	if err != nil {
		return GetLiveOrdersResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetLiveOrdersResult{}, err
	}

	defer resp.Body.Close()

	reader, err := getReadCloserFromResponse(resp)
	if err != nil {
		return GetLiveOrdersResult{}, err
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetLiveOrdersResult{}, err
	}

	v := GetLiveOrdersResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetLiveOrdersResult{}, err
	}

	if v.Error.Code != "" {
		return v, errors.New(v.Error.Message)
	}

	return v, nil
}
