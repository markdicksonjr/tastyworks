package tastyworks

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type GetAccountsResult struct {
	Data GetAccountsResultData
}

type GetAccountsResultData struct {
	Items []GetAccountsResultItem
}

type GetAccountsResultItem struct {
	Account        GetAccountsResultAccount
	AuthorityLevel string `json:"authority-level"`
}

type GetAccountsResultAccount struct {
	AccountNumber        string `json:"account-number"`
	ExternalId           string `json:"external-id"`
	OpenedAt             string `json:"opened-at"`
	Nickname             string
	AccountTypeName      string `json:"account-type-name"`
	DayTraderStatus      bool   `json:"day-trader-status"`
	IsClosed             bool   `json:"is-closed"`
	IsFirmError          bool   `json:"is-firm-error"`
	IsFirmProprietary    bool   `json:"is-firm-proprietary"`
	IsTestDrive          bool   `json:"is-test-drive"`
	MarginOrCash         string `json:"margin-or-cash"`
	IsForeign            bool   `json:"is-foreign"`
	FundingDate          string `json:"funding-date"`
	InvestmentObjective  string `json:"investment-objective"`
	SuitableOptionsLevel string `json:"suitable-options-level"`
	CreatedAt            string `json:"created-at"`
}

func GetAccounts(sessionToken string) (GetAccountsResult, error) {
	req, err := http.NewRequest("GET", TastyTradeHost+"/customers/me/accounts", nil)

	if err != nil {
		return GetAccountsResult{}, err
	}

	SetStandardHeadersOnRequest(req)

	req.Header.Set("Authorization", sessionToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return GetAccountsResult{}, err
	}

	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return GetAccountsResult{}, err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return GetAccountsResult{}, err
	}

	v := GetAccountsResult{}
	if err := json.Unmarshal(body, &v); err != nil {
		return GetAccountsResult{}, err
	}

	return v, nil
}
