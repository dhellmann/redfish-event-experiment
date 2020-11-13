package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type BMCSettings struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ReceiverSettings struct {
	Endpoint string `yaml:"endpoint"`
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
}

// Settings includes all of the application settinggs
type Settings struct {
	BMC      BMCSettings      `yaml:"bmc"`
	Receiver ReceiverSettings `yaml:"receiver"`
}

// LoadFromFile reads the named file and returns the Settings
// contained within.
func LoadFromFile(filename string) (*Settings, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := Settings{}
	err = yaml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
