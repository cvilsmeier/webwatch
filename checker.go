package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// -----------------

type CheckResult struct {
	Ok   bool
	Text string
}

func (this CheckResult) IsDifferent(other CheckResult) bool {
	return this.Ok != other.Ok || this.Text != other.Text
}

// -----------------

type Checker interface {
	Check() CheckResult
}

func NewChecker(urls []string) Checker {
	return checkerImpl{
		urls,
	}
}

type checkerImpl struct {
	urls []string
}

func (this checkerImpl) Check() CheckResult {
	ok := true
	text := ""
	for _, url := range this.urls {
		url = strings.TrimSpace(url)
		if url != "" {
			err := this.checkUrl(url)
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

func (this checkerImpl) checkUrl(url string) error {
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
