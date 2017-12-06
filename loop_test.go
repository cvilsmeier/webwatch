package webwatch

import (
	"fmt"
	"testing"
	"time"
)

// ------------------------------------------

var _ Sleeper = &fakeSleeper{}

type fakeSleeper struct {
	calls []string
}

func (fake *fakeSleeper) Sleep(duration time.Duration) {
	call := fmt.Sprintf("Sleep(%s)", duration)
	fake.calls = append(fake.calls, call)
}

func (fake *fakeSleeper) reset() {
	fake.calls = []string{}
}

// ------------------------------------------

var _ Checker = &fakeChecker{}

type fakeChecker struct {
	result CheckResult
}

func (fake *fakeChecker) Check() CheckResult {
	return fake.result
}

// ------------------------------------------

var _ Mailer = &fakeMailer{}

type fakeMailer struct {
	calls []string
}

func (fake *fakeMailer) SendRestarted(now time.Time) {
	call := fmt.Sprintf("SendRestarted()")
	fake.calls = append(fake.calls, call)
}

func (fake *fakeMailer) SendReport(now time.Time, ok bool, text string) {
	call := fmt.Sprintf("SendReport(%t, %s)", ok, text)
	fake.calls = append(fake.calls, call)
}

func (fake *fakeMailer) reset() {
	fake.calls = []string{}
}

// ------------------------------------------

func TestLoop(t *testing.T) {
	sleeper := &fakeSleeper{[]string{}}
	checker := &fakeChecker{CheckResult{true, "ok"}}
	mailer := &fakeMailer{[]string{}}
	loop := NewLoop(
		sleeper,
		checker,
		mailer,
		mustParseDuration("5m"),  // checks,
		mustParseDuration("12h"), // reports,
		mustParseDuration("1h"),  // limit
	)
	now := time.Date(2017, 1, 1, 10, 0, 0, 0, time.UTC)
	// first tick
	// -> send report
	// -> sleep limit
	loop.Tick(now)
	assertEqInt(t, 1, len(mailer.calls))
	assertEqStr(t, "SendReport(true, ok)", mailer.calls[0])
	mailer.reset()
	assertEqInt(t, 1, len(sleeper.calls))
	assertEqStr(t, "Sleep(1h0m0s)", sleeper.calls[0])
	sleeper.reset()
	// after 1h: new check
	// -> no report
	// -> sleep checks (5m)
	now = now.Add(1 * time.Hour)
	loop.Tick(now)
	assertEqInt(t, 0, len(mailer.calls))
	mailer.reset()
	assertEqInt(t, 1, len(sleeper.calls))
	assertEqStr(t, "Sleep(5m0s)", sleeper.calls[0])
	sleeper.reset()
	// after 5m: new check, url error
	// -> send report
	// -> sleep limit (1h)
	now = now.Add(5 * time.Minute)
	checker.result = CheckResult{false, "err"}
	loop.Tick(now)
	assertEqInt(t, 1, len(mailer.calls))
	assertEqStr(t, "SendReport(false, err)", mailer.calls[0])
	mailer.reset()
	assertEqInt(t, 1, len(sleeper.calls))
	assertEqStr(t, "Sleep(1h0m0s)", sleeper.calls[0])
	sleeper.reset()
	// after 1h: new check
	// -> no report
	// -> sleep checks (5m)
	now = now.Add(1 * time.Hour)
	loop.Tick(now)
	assertEqInt(t, 0, len(mailer.calls))
	mailer.reset()
	assertEqInt(t, 1, len(sleeper.calls))
	assertEqStr(t, "Sleep(5m0s)", sleeper.calls[0])
	sleeper.reset()
	// after 12h: new check
	// -> report after 12h
	// -> sleep checks (5m)
	now = now.Add(12 * time.Hour)
	loop.Tick(now)
	assertEqInt(t, 1, len(mailer.calls))
	assertEqStr(t, "SendReport(false, err)", mailer.calls[0])
	mailer.reset()
	assertEqInt(t, 1, len(sleeper.calls))
	assertEqStr(t, "Sleep(1h0m0s)", sleeper.calls[0])
	sleeper.reset()
}

// ------------------------------------------

func mustParseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}
