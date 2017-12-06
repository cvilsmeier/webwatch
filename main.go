package main

import (
	"flag"
	"os"
	"time"
)

var configfile string
var verbose bool

func main() {
	// cli args
	flag.StringVar(&configfile, "config", "config.json", "the name of the config file")
	flag.BoolVar(&verbose, "v", false, "verbose output (default off)")
	flag.Parse()
	// logging
	if verbose {
		DEBUG.SetOutput(os.Stdout)
	}
	// config
	config, err := LoadConfig(configfile)
	if err != nil {
		INFO.Fatal(err)
	}
	// setup
	sleeper := NewSleeper()
	checker := NewChecker(config.Urls)
	mailer := NewMailer(config.Mail)
	loop := NewLoop(sleeper, checker, mailer, time.Duration(config.Checks), time.Duration(config.Reports), time.Duration(config.Limit))
	// loop
	mailer.SendRestarted(time.Now())
	for {
		loop.Tick(time.Now())
	}
}
