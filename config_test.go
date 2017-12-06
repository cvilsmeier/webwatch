package main

import (
	"testing"
)

func TestConfigSample(t *testing.T) {
	c, err := LoadConfig("testdata/config_sample.json")
	assertNil(t, err)
	assertEqInt(t, 2, len(c.Urls))
	assertEqStr(t, "https://www.google.com", c.Urls[0])
	assertEqStr(t, "https://www.twitter.com", c.Urls[1])
	assertEqStr(t, "5m0s", c.Checks.String())
	assertEqStr(t, "12h0m0s", c.Reports.String())
	assertEqStr(t, "1h0m0s", c.Limit.String())
	assertEqStr(t, "[Webwatch] MY_SERVER", c.Mail.Subject)
	assertEqStr(t, "mail@example.com", c.Mail.From)
	assertEqStr(t, "myself@example.com", c.Mail.To)
	assertEqStr(t, "smtp.example.com", c.Mail.Host)
	assertEqStr(t, "example_username_00012", c.Mail.Username)
	assertEqStr(t, "example_password_00012", c.Mail.Password)
}

func TestConfigEmpty(t *testing.T) {
	c, err := LoadConfig("testdata/config_empty.json")
	assertNil(t, err)
	assertEqInt(t, 0, len(c.Urls))
	assertEqStr(t, "1s", c.Checks.String())
	assertEqStr(t, "30s", c.Reports.String())
	assertEqStr(t, "30s", c.Limit.String())
	assertEqStr(t, "Webwatch", c.Mail.Subject)
	assertEqStr(t, "", c.Mail.From)
	assertEqStr(t, "", c.Mail.To)
	assertEqStr(t, "", c.Mail.Host)
	assertEqStr(t, "", c.Mail.Username)
	assertEqStr(t, "", c.Mail.Password)
}
