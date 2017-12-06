package main

import (
	"flag"
	"github.com/cvilsmeier/webwatch"
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
		webwatch.DEBUG.SetOutput(os.Stdout)
	}
	// config
	webwatch.DEBUG.Printf("loading %q", configfile)
	config, err := webwatch.LoadConfig(configfile)
	if err != nil {
		webwatch.INFO.Fatal(err)
	}
	// setup
	sleeper := webwatch.NewSleeper()
	checker := webwatch.NewChecker(config.Urls)
	mailer := webwatch.NewMailer(config.Mail)
	loop := webwatch.NewLoop(sleeper, checker, mailer, time.Duration(config.Checks), time.Duration(config.Reports), time.Duration(config.Limit))
	// loop
	mailer.SendRestarted(time.Now())
	for {
		loop.Tick(time.Now())
	}
}
