package webwatch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// A CheckResult contains the result of checking a number
// of URLs for their reachability.
type CheckResult struct {
	Ok   bool
	Text string
}

// IsDifferent returns true if cr differs from other.
func (cr CheckResult) IsDifferent(other CheckResult) bool {
	return cr.Ok != other.Ok || cr.Text != other.Text
}

// -----------------

// A Checker checks URLs and returns a CheckResult.
type Checker interface {

	// Check checks URLs and returns a CheckResult.
	Check() CheckResult
}

// NewChecker creates a new Checker that checks a list of URLs.
func NewChecker(urls []string) Checker {
	return checkerImpl{
		urls,
	}
}

// -----------------

type checkerImpl struct {
	urls []string
}

func (ci checkerImpl) Check() CheckResult {
	ok := true
	text := ""
	for _, url := range ci.urls {
		url = strings.TrimSpace(url)
		if url != "" {
			err := ci.checkURL(url)
			if err != nil {
				ok = false
				text += fmt.Sprintf("%s\r\nERR %s\r\n\r\n", url, err)
			} else {
				text += fmt.Sprintf("%s\r\nOK\r\n\r\n", url)
			}
		}
	}
	return CheckResult{ok, text}
}

func (ci checkerImpl) checkURL(url string) error {
	DEBUG.Printf("check %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		DEBUG.Printf(" ERR %s\n", err)
		return err
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		DEBUG.Printf(" ERR %s\n", err)
		return err
	}
	if res.StatusCode != 200 {
		err = fmt.Errorf("status %d", res.StatusCode)
		DEBUG.Printf(" ERR %s\n", err)
		return err
	}
	DEBUG.Printf(" OK\n")
	return nil
}
