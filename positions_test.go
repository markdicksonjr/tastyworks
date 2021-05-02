package tastyworks

import (
	"os"
	"testing"
)

func TestGetPositions(t *testing.T) {
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
	pos, err := GetPositions(v.Data.SessionToken, acc.Data.Items[0].Account.AccountNumber)
	if err != nil {
		t.Fatal(err)
	}

	// NOTE: obviously, fails if the account has no active positions
	if len(pos.Data.Items) == 0 {
		t.Fatal("empty position list")
	}
}
