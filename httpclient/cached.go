package httpclient

import (
	"fmt"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func getCache() httpcache.Cache {
	cache := httpcache.NewMemoryCache()

	dir, err := os.UserCacheDir()
	if err != nil {
		return cache
	}

	dir = filepath.Join(dir, "gh-notifications")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return cache
	}

	return diskcache.New(dir)
}

func NewCachedOauthClient(maxAge time.Duration, ts oauth2.TokenSource)* http.Client {
	cache := getCache()

	oat := &oauth2.Transport{
		Source: ts,
	}

	ct := &httpcache.Transport{Transport: oat, Cache: cache}

	lt := liberalRoundTripper{
		rt:  ct,
		max: time.Minute * 10,
	}

	return &http.Client{
		Transport: lt,
	}
}

// liberalRoundTripper is a round-tripper that allows requests to go stale.
type liberalRoundTripper struct {
	rt http.RoundTripper
	max time.Duration
}

// Implements RoundTrip
func (l liberalRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	cc := fmt.Sprintf("max-stale=%d", int(l.max / time.Second))

	if existing := request.Header.Get("cache-control"); existing != "" {
		cc = fmt.Sprintf("%s,%s", existing, cc)
	}

	request.Header.Add("cache-control", cc)
	return l.rt.RoundTrip(request)
}
