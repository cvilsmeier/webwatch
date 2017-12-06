package webwatch

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	// INFO logger, writes to stdout.
	INFO  *log.Logger = log.New(os.Stdout, "INFO ", log.LstdFlags)

	// DEBUG logger, writes to stdout. Default disabled.
	DEBUG *log.Logger = log.New(ioutil.Discard, "DEBG ", log.LstdFlags)
)
