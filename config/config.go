package config

import "github.com/hashicorp/hcl/v2/hclsimple"

// Config is the top-level config
type Config struct {
	Server Server `hcl:"server,block"`
	Hpotato Hpotato `hcl:"hpotato,block"`
}

// Server is gRPC server related config
type Server struct {
	BindAddr string `hcl:"bind_addr,attr"`
}

// GetConfig loads config from an hcl file at the specified path
func GetConfig(path string) (Config, error) {
	config := Config{}
	err := hclsimple.DecodeFile(path, nil, &config)

	return config, err
}

// DefaultConfig returns default config
func DefaultConfig() Config {
	return Config{
		Server: Server{
			BindAddr: ":8080",
		},
		Hpotato: DefaultHpotatoConfig(),
	}
}
