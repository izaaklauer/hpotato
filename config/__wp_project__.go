package config

type %%Wp_project%% struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// Default%%Wp_project%%Config returns default config values
func Default%%Wp_project%%Config() %%Wp_project%% {
	return %%Wp_project%%{
	HelloWorldMessage:
		"hello from the default config",
	}
}
