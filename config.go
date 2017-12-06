package webwatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// --------------------------------------------

// Interval is a time.Duration that can be marshalled to
// and unmarshalled from JSON.
type Interval time.Duration

// UnmarshalJSON converts an Interval to JSON.
func (iv *Interval) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		text := string(b[1 : len(b)-1])
		d, err := time.ParseDuration(text)
		if err != nil {
			return err
		}
		*iv = Interval(d)
		return nil
	}
	return fmt.Errorf("does not begin with \"")
}

// MarshalJSON converts a JSON value to an Interval.
func (iv Interval) MarshalJSON() ([]byte, error) {
	tstamp := iv.String()
	quoted := fmt.Sprintf("\"%s\"", tstamp)
	return []byte(quoted), nil
}

// String() returns a string representation of this Interval,
// see https://golang.org/pkg/time/#Duration.String
// for more info.
func (iv Interval) String() string {
	return time.Duration(iv).String()
}

// --------------------------------------------

// Config holds the configuration data
// found in the config.json file.
type Config struct {
	Urls    []string   `json:"urls"`
	Checks  Interval   `json:"checks"`
	Reports Interval   `json:"reports"`
	Limit   Interval   `json:"limit"`
	Mail    MailConfig `json:"mail"`
}

// MailConfig holds the mail configuration
// found in the config.json file.
type MailConfig struct {
	Subject  string `json:"subject"`
	From     string `json:"from"`
	To       string `json:"to"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoadConfig loads the JSON configuration file and
// converts it to a Config structure.
// LoadConfig returns an error if the fil cannot be loaded
// or parsed.
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
