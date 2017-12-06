package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// --------------------------------------------

type Interval time.Duration

func (this *Interval) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		text := string(b[1 : len(b)-1])
		d, err := time.ParseDuration(text)
		if err != nil {
			return err
		}
		*this = Interval(d)
		return nil
	}
	return fmt.Errorf("does not begin with \"")
}

func (this Interval) MarshalJSON() ([]byte, error) {
	tstamp := this.String()
	quoted := fmt.Sprintf("\"%s\"", tstamp)
	return []byte(quoted), nil
}

func (this Interval) String() string {
	return time.Duration(this).String()
}

// --------------------------------------------

type Config struct {
	Urls    []string   `json:"urls"`
	Checks  Interval   `json:"checks"`
	Reports Interval   `json:"reports"`
	Limit   Interval   `json:"limit"`
	Mail    MailConfig `json:"mail"`
}

type MailConfig struct {
	Subject  string `json:"subject"`
	From     string `json:"from"`
	To       string `json:"to"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadConfig(filename string) (Config, error) {
	c := Config{}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(buf, &c)
	if err != nil {
		return c, err
	}
	c = sanitizeConfig(c)
	return c, nil
}

func sanitizeConfig(c Config) Config {
	if time.Duration(c.Checks) < 1*time.Second {
		c.Checks = Interval(1 * time.Second)
	}
	if time.Duration(c.Reports) < 30*time.Second {
		c.Reports = Interval(30 * time.Second)
	}
	if time.Duration(c.Limit) < 30*time.Second {
		c.Limit = Interval(30 * time.Second)
	}
	if c.Mail.Subject == "" {
		c.Mail.Subject = "Webwatch"
	}
	return c
}
