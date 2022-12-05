package config

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	require := require.New(t)

	t.Run("Can render local config", func(t *testing.T) {
		config := Config{}
		err := hclsimple.DecodeFile("config_local.hcl", nil, &config)
		require.NoError(err)
		require.NotEmpty(config.Server.BindAddr)

		// Add more app-specific config property checks here if desired.
	})
}
