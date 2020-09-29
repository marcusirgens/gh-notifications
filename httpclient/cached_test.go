package httpclient

import (
	"net/http"
	"testing"
)

// This is an interface guard. This test will fail to compie if the liberalRoundTripper
// type stops implementing http.RoundTripper.
func TestLiberalRoundTripper_ImplementsRoundTripper(t *testing.T) {
	var _ http.RoundTripper = liberalRoundTripper{}
}