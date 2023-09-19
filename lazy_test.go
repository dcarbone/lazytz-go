package lazytz_test

import (
	"testing"

	"github.com/dcarbone/lazytz-go"
)

func TestCanDoThing(t *testing.T) {
	for _, abbr := range lazytz.GetAbbreviations() {
		_, ok := lazytz.Get(abbr)
		if !ok {
			t.Logf("Unable to load info for %q", abbr)
			t.Fail()
		}
	}
}
