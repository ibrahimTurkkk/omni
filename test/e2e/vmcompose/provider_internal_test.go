package vmcompose

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// SetupDataFixtures returns the test fixture filenames of manifest and infrastructure data files.
// This accesses private types, so it's in the same package as the types.
func SetupDataFixtures(t *testing.T) (string, string) {
	t.Helper()

	// Write manifest to disk
	manifest := `
network = "devnet"
anvil_chains = ["chain_a"]
multi_omni_evms = true

[node.validator01]
[node.validator02]
`
	manifestFile := filepath.Join(t.TempDir(), "test.toml")
	err := os.WriteFile(manifestFile, []byte(manifest), 0o644)
	require.NoError(t, err)

	const vm1, vm2, vm3 = "vm1", "vm2", "vm3"

	dataJSON := dataJSON{
		NetworkCIDR: "127.0.0.1/24",
		VMs: []vmJSON{
			{Name: vm1, IP: "127.0.0.1"},
			{Name: vm2, IP: "127.0.0.2"},
			{Name: vm3, IP: "127.0.0.3"},
		},
		ServicesByVM: map[string]string{
			"validator01":     vm1,
			"validator01_evm": vm1,

			"validator02":     vm2,
			"validator02_evm": vm2,

			"chain_a": vm3,
			"relayer": vm3,
		},
	}

	// Write raw data json to disk
	bz, err := json.Marshal(dataJSON)
	require.NoError(t, err)
	dataFile := filepath.Join(t.TempDir(), "data.json")
	err = os.WriteFile(dataFile, bz, 0o644)
	require.NoError(t, err)

	return manifestFile, dataFile
}
