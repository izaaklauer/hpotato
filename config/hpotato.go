package config

type Hpotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultHpotatoConfig returns default config values
func DefaultHpotatoConfig() Hpotato {
	return Hpotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
