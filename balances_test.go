package tastyworks

import (
	"os"
	"testing"
)

func TestGetAccountBalance(t *testing.T) {
	v, err := Authorize(os.Getenv("TT_USER"), os.Getenv("TT_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	acc, err := GetAccounts(v.Data.SessionToken)
	if err != nil {
		t.Fatal(err)
	}
	if len(acc.Data.Items) == 0 {
		t.Fatal("empty account result")
	}
	bal, err := GetAccountBalance(v.Data.SessionToken, acc.Data.Items[0].Account.AccountNumber)
	_ = bal
}
