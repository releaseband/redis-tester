package options

import "time"

type TimeOuts struct {
	Read  time.Duration `yaml:"read"`
	Write time.Duration `yaml:"write"`
	Dial  time.Duration `yaml:"dial"`
	Pool  time.Duration `yaml:"pool"`
	Idle  time.Duration `yaml:"idle"`
}

type Retries struct {
	MaxRetries      int           `yaml:"max_retries"`
	MinRetryBackoff time.Duration `yaml:"min_retry"`
	MaxRetryBackoff time.Duration `yaml:"max_retry"`
}

type Auth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type General struct {
	Auth               `yaml:"auth"`
	TimeOuts           `yaml:"timeouts"`
	Retries            `yaml:"retries"`
	MinIdleConns       int           `yaml:"min_idle_conns"`
	PoolSize           int           `yaml:"pool_size"`
	MaxConnAge         time.Duration `yaml:"max_conn_age"`
	IdleCheckFrequency time.Duration `yaml:"idle_check_frequency"`
}

type Redis struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type RedisSettings struct {
	General
	Redis
}

type ClusterSettings struct {
	General
	Cluster
}

type TesterSettings struct {
	General General `yaml:"general"`
	Redis   Redis   `yaml:"redis"`
	Cluster Cluster `yaml:"cluster"`
	Test    Test    `yaml:"test"`
}

func (ts TesterSettings) RedisSettings() RedisSettings {
	return RedisSettings{
		General: ts.General,
		Redis:   ts.Redis,
	}
}

func (ts TesterSettings) ClusterSettings() ClusterSettings {
	return ClusterSettings{
		General: ts.General,
		Cluster: ts.Cluster,
	}
}
