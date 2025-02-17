package netmap

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCase represents collection of placement policy tests for a single node set.
type TestCase struct {
	Name  string     `json:"name"`
	Nodes []NodeInfo `json:"nodes"`
	Tests map[string]struct {
		Policy    PlacementPolicy `json:"policy"`
		Pivot     []byte          `json:"pivot,omitempty"`
		Result    [][]int         `json:"result,omitempty"`
		Error     string          `json:"error,omitempty"`
		Placement struct {
			Pivot  []byte
			Result [][]int
		} `json:"placement,omitempty"`
	}
}

func compareNodes(t *testing.T, expected [][]int, nodes Nodes, actual []Nodes) {
	require.Equal(t, len(expected), len(actual))
	for i := range expected {
		require.Equal(t, len(expected[i]), len(actual[i]))
		for j, index := range expected[i] {
			require.Equal(t, nodes[index], actual[i][j])
		}
	}
}

func TestPlacementPolicy_Interopability(t *testing.T) {
	const testsDir = "./json_tests"

	f, err := os.Open(testsDir)
	require.NoError(t, err)

	ds, err := f.ReadDir(0)
	require.NoError(t, err)

	for i := range ds {
		bs, err := ioutil.ReadFile(filepath.Join(testsDir, ds[i].Name()))
		require.NoError(t, err)

		var tc TestCase
		require.NoError(t, json.Unmarshal(bs, &tc), "cannot unmarshal %s", ds[i].Name())

		t.Run(tc.Name, func(t *testing.T) {
			nodes := NodesFromInfo(tc.Nodes)
			nm, err := NewNetmap(nodes)
			require.NoError(t, err)

			for name, tt := range tc.Tests {
				t.Run(name, func(t *testing.T) {
					v, err := nm.GetContainerNodes(&tt.Policy, tt.Pivot)
					if tt.Result == nil {
						require.Error(t, err)
						require.Contains(t, err.Error(), tt.Error)
					} else {
						require.NoError(t, err)

						res := v.Replicas()
						compareNodes(t, tt.Result, nodes, res)

						if tt.Placement.Result != nil {
							res, err := nm.GetPlacementVectors(v, tt.Placement.Pivot)
							require.NoError(t, err)
							compareNodes(t, tt.Placement.Result, nodes, res)
						}
					}
				})
			}
		})
	}
}
func BenchmarkManySelects(b *testing.B) {
	testsFile := filepath.Join("json_tests", "many_selects.json")
	bs, err := ioutil.ReadFile(testsFile)
	require.NoError(b, err)

	var tc TestCase
	require.NoError(b, json.Unmarshal(bs, &tc))
	tt, ok := tc.Tests["Select"]
	require.True(b, ok)

	nodes := NodesFromInfo(tc.Nodes)
	nm, err := NewNetmap(nodes)
	require.NoError(b, err)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err = nm.GetContainerNodes(&tt.Policy, tt.Pivot)
		if err != nil {
			b.FailNow()
		}
	}
}
