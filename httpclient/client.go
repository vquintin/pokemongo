package pokemongo

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
)

var ErrSentinel = errors.New("Sentinel error")

func NewClient() *http.Client {
	checkRedirect := func(req *http.Request, via []*http.Request) error {
		return ErrSentinel
	}
	cookieJar, _ := cookiejar.New(nil)
	return &http.Client{CheckRedirect: checkRedirect, Jar: cookieJar}
}
