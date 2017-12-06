package webwatch

import (
	"time"
)

// Loop is the main loop.
type Loop struct {
	sleeper    Sleeper
	checker    Checker
	mailer     Mailer
	checks     time.Duration
	reports    time.Duration
	limit      time.Duration
	lastResult CheckResult
	lastMail   time.Time
}

// NewLoop creates a new Loop.
func NewLoop(sleeper Sleeper, checker Checker, mailer Mailer, checks, reports, limit time.Duration) *Loop {
	return &Loop{
		sleeper,
		checker,
		mailer,
		checks,
		reports,
		limit,
		CheckResult{},
		time.Time{},
	}
}

// Tick executes the main loop.
// It invokes the checker and (maybe) sends a mail and
// then sleeps for a certain time.
func (l *Loop) Tick(now time.Time) {
	result := l.checker.Check()
	if result.IsDifferent(l.lastResult) {
		l.mailer.SendReport(now, result.Ok, result.Text)
		l.lastMail = now
		l.lastResult = result
		l.sleeper.Sleep(l.limit)
	} else if now.Sub(l.lastMail) > l.reports {
		l.mailer.SendReport(now, l.lastResult.Ok, l.lastResult.Text)
		l.lastMail = now
		l.sleeper.Sleep(l.limit)
	} else {
		l.sleeper.Sleep(l.checks)
	}
}
