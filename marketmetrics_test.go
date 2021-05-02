package tastyworks

import (
	"os"
	"testing"
)

func TestGetMarketMetricsForTickers(t *testing.T) {
	v, err := Authorize(os.Getenv("TT_USER"), os.Getenv("TT_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	metrics, err := GetMarketMetricsForTickers(v.Data.SessionToken, []string{"SPY","QQQ"})
	if err != nil {
		t.Fatal(err)
	}

	if len(metrics.Data.Items) != 2 {
		t.Fatal("did not get correct number of metrics results - expected 2, got ", len(metrics.Data.Items))
	}

	if metrics.Data.Items[0].Symbol == "" {
		t.Fatal("empty metrics result first item")
	}
}
