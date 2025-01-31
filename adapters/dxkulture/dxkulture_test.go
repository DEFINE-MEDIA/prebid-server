package dxkulture

import (
	"testing"

	"github.com/prebid/prebid-server/config"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder("dxkulture", config.Adapter{Endpoint: "https://ads.kulture.media/pbs"}, config.Server{ExternalUrl: "http://hosturl.com", GvlID: 1, DataCenter: "2"})
	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}
	adapterstest.RunJSONBidderTest(t, "dxkulturetest", bidder)
}
