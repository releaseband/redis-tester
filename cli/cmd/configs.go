package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ayupov-ayaz/redis-tester/options"
	"gopkg.in/yaml.v3"
)

func makeTesterSettings(cfgFileName string) (*options.TesterSettings, error) {
	if _, err := os.Stat(cfgFileName); err != nil {
		return nil, fmt.Errorf("check file `%s` failed: %w", cfgFileName, err)
	}

	data, err := ioutil.ReadFile(cfgFileName)
	if err != nil {
		return nil, fmt.Errorf("read '%s' file failed: %w", cfgFileName, err)
	}

	ts := &options.TesterSettings{}
	if err := yaml.Unmarshal(data, ts); err != nil {
		return nil, fmt.Errorf("unmarshal configs failed: %w", err)
	}

	return ts, nil
}
