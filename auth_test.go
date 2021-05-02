package tastyworks

import (
	"os"
	"testing"
)

func TestAuthorize(t *testing.T) {
	v, err := Authorize(os.Getenv("TT_USER"), os.Getenv("TT_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	if v.Data.User.ExternalId == "" {
		t.Fatal("no external ID in response")
	}
}
