package main

import (
	"io/ioutil"
	"log"
	"os"
)

var INFO *log.Logger = log.New(os.Stdout, "INFO  ", log.LstdFlags)
var DEBUG *log.Logger = log.New(ioutil.Discard, "DEBUG ", log.LstdFlags)
