package options

type Slot struct {
	Start     int      `yaml:"start"`
	End       int      `yaml:"end"`
	Addresses []string `yaml:"addresses"`
}

type ManualCluster struct {
	Use   bool   `yaml:"use"`
	Slots []Slot `yaml:"slots"`
}

type Cluster struct {
	Use            bool          `yaml:"use"`
	Manual         ManualCluster `yaml:"manual"`
	Addresses      []string      `yaml:"addresses"`
	ReadOnly       bool          `yaml:"read_only"`
	MaxRedirects   int           `yaml:"max_redirects"`
	RouteRandomly  bool          `yaml:"route_randomly"`
	RouteByLatency bool          `yaml:"route_by_latency"`
}
