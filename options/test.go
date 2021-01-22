package options

import "time"

type Commands struct {
	Set    bool `yaml:"set"`
	Get    bool `yaml:"get"`
	Del    bool `yaml:"del"`
	RPush  bool `yaml:"r_push"`
	LRange bool `yaml:"l_range"`
	LTrim  bool `yaml:"l_trim"`
}

type Test struct {
	Iterations     int           `yaml:"iterations"`
	SetExpirations time.Duration `yaml:"expirations"`
	Commands       Commands      `yaml:"commands"`
}

func (t Test) Count() int {
	return t.Iterations
}

func (t Test) Expiration() time.Duration {
	return t.SetExpirations
}
