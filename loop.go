package main

import (
	"time"
)

// Loop ist the main loop.
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

// Tick executes the main loop. It invokes the checker and (maybe) sends a mail and sleeps for a certain time.
func (this *Loop) Tick(now time.Time) {
	result := this.checker.Check()
	if result.IsDifferent(this.lastResult) {
		this.mailer.SendReport(now, result.Ok, result.Text)
		this.lastMail = now
		this.lastResult = result
		this.sleeper.Sleep(this.limit)
	} else if now.Sub(this.lastMail) > this.reports {
		this.mailer.SendReport(now, this.lastResult.Ok, this.lastResult.Text)
		this.lastMail = now
		this.sleeper.Sleep(this.limit)
	} else {
		this.sleeper.Sleep(this.checks)
	}
}
