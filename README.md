# Webwatch

[![GoDoc](https://godoc.org/github.com/cvilsmeier/webwatch?status.svg)](https://godoc.org/github.com/cvilsmeier/webwatch)
[![Build Status](https://travis-ci.org/cvilsmeier/webwatch.svg?branch=master)](https://travis-ci.org/cvilsmeier/webwatch)
[![Go Report Card](https://goreportcard.com/badge/github.com/cvilsmeier/webwatch)](https://goreportcard.com/report/github.com/cvilsmeier/webwatch)

**Webwatch checks HTTP/HTTPS URLs and sends mail if there's a problem.**

You run Webwatch on your server(s) and let it watch local or remote URLs. When
the state of one or more URLS change (from 'reachable' to 'not reachable' or
vice versa), Webwatch will send an email. You can rate-limit emails to avoid
being flooded.

If you run two or more servers, install Webwatch on each of them and let them
check each other.


## Installation

Webwatch is written in [Go](https://golang.org/) (we need 1.9.0 or higher).

```bash
$ go get -u github.com/cvilsmeier/webwatch
```

## Usage

```bash
$ webwatch -help
```
shows a help page.


```bash
$ webwatch -config home/cv/webwatch/config.json
```
runs Webwatch with a config file.


## Configuration

Webwatch loads its configured from a json file with the following structure:

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
  'm' for minutes and 's' for seconds. Minimum is '1s'.

* `reports` The interval Webwatch should send report mails. This setting
  applies only if the states the URLs did not change since the last mail.  If
  the state of one URL changes, Webwatch will send mail immediately and not
  wait for the 'reports' interval.

* `limit` Rate-limit for mails. Webwatch will sent at most one mail per 'limit'
  period

* `mail` Mail configuration.

    * `subject` The mail subject. Webwatch will append "restarted" or
      "OK" or "ERR" to that subject, depending on the state of your URLs.

    * `from` The from address.

    * `to` The to address.

    * `host` The host name of the smtp server Webwatch sends mail to.
      Webwatch will try to talk smtp over port 25.

    * `username` The username for authenticating against the smtp server.

    * `password` The password for that username.


## Author

Me, C. Vilsmeier


## License

[![wtfpl](http://www.wtfpl.net/wp-content/uploads/2012/12/wtfpl-badge-1.png)](http://www.wtfpl.net)

