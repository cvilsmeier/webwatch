# Webwatch

[![GoDoc](https://godoc.org/github.com/cvilsmeier/webwatch?status.svg)](https://godoc.org/github.com/cvilsmeier/webwatch)
[![Build Status](https://travis-ci.org/cvilsmeier/webwatch.svg?branch=master)](https://travis-ci.org/cvilsmeier/webwatch)
[![Go Report Card](https://goreportcard.com/badge/github.com/cvilsmeier/webwatch)](https://goreportcard.com/report/github.com/cvilsmeier/webwatch)

**Webwatch checks HTTP/HTTPS URLs and sends mail if there's a problem.**

You run webwatch on your server(s) and let it watch local or remote URLs. When
the state of one or more URLS change (from 'reachable' to 'not reachable' or
vice versa), webwatch will send an email. You can rate-limit emails to avoid
being flooded.

If you run two or more servers, install webwatch on each of them and let them
check each other.


## Build

Webwatch is written in [Go](https://golang.org/) (we need 1.9.0 or higher).

```bash
user@wombat # mkdir ~/webwatch
user@wombat # export GOPATH=/tmp/webwatch
user@wombat # cd $GOPATH
user@wombat # go get -u github.com/cvilsmeier/webwatch/cmd/webwatch
user@wombat # chmod 700 bin/webwatch
user@wombat # bin/webwatch -help
```

## Usage

```bash
user@wombat ~ # webwatch -help
Usage of webwatch:
  -config string
        the name of the config file (default "config.json")
  -v    verbose output (default off)
```

```bash
user@wombat ~ # webwatch -config home/cv/webwatch/config.json -v
```


## Configuration

Webwatch loads its configuration from a json file with the following structure:

```json
{
    "urls": [
        "https://www.google.com",
        "https://www.twitter.com"
    ],
    "checks": "5m",
    "reports": "12h",
    "limit": "1h",
    "mail" : {
        "subject" : "[Webwatch] MY_SERVER",
        "from" : "mail@example.com",
        "to" : "myself@example.com",
        "host" : "smtp.example.com",
        "username" : "example_username_00012",
        "password" : "example_password_00012"
    }
}
```

* `urls` The URLs you want to watch. Configure any number of URLs here.

* `checks` The check interval for your URLs. Valid suffixes are 'h' for hours,
  'm' for minutes and 's' for seconds. Default is '5m'.

* `reports` The interval webwatch should send report mails. This setting
  applies only if the states the URLs did not change since the last mail.  If
  the state of one URL changes, webwatch will send mail immediately and not
  wait for the 'reports' interval. Default is '12h'.

* `limit` Rate-limit for mails. webwatch will sent at most one mail per 'limit'
  period, no matter what. Default is '1h'.

* `mail` Mail configuration.

    * `subject` The mail subject. webwatch will append "restarted" or
      "OK" or "ERR" to that subject, depending on the state of your URLs.

    * `from` The from address.

    * `to` The to address.

    * `host` The host name of the smtp server webwatch sends mail to.
      webwatch will try to talk smtp over port 25.

    * `username` The username for authenticating against the smtp server.

    * `password` The password for that username.


## Author

C. Vilsmeier


## License

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>

