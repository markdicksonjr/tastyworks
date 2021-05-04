package tastyworks

import (
	"os"
	"testing"
)

func TestGetQuoteStreamerToken(t *testing.T) {
	v, err := Authorize(os.Getenv("TT_USER"), os.Getenv("TT_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	pos, err := GetQuoteStreamerToken(v.Data.SessionToken)
	if err != nil {
		t.Fatal(err)
	}

	_ = pos

	//// NOTE: tests fail if account has no transactions
	//if len(pos.Data.Items) == 0 {
	//	t.Fatal("empty transaction list")
	//}
}
