package configurator

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func yamlFromFile(t *testing.T) []byte {
	file, err := ioutil.ReadFile("example/all-envs.yaml")
	require.NoError(t, err)
	require.NotNil(t, file)

	return file
}
func TestNewConfigFromYaml(t *testing.T) {

	result := NewConfigFromYaml(yamlFromFile(t))

	require.Equal(t, "foo", result.Name, "name did not match")
}

func TestOutputsADefaultConfigMap(t *testing.T) {
	cfg := NewConfigFromYaml(yamlFromFile(t))

	all := cfg.OutputAll()

	require.True(t, len(all) > 0, "results did not contain any cm's")

	result, exists := all["default"]
	require.True(t, exists, "default did not exist in results")
	require.Equal(t, "foo", result.Name, "name did not match")
	require.Equal(t, "system", result.Namespace, "namespace did not fatch")
	require.Equal(t, "value1", result.Data["setting1"], "setting1 did not match")
	require.Equal(t, "value2", result.Data["setting2"], "setting1 did not match")
	require.Equal(t, "value1", result.Annotations["key1"], "annotation1 did not match")
	require.Equal(t, "value2", result.Annotations["key2"], "annotation2 did not match")

}

func TestOutputsAnOverride(t *testing.T) {
	cfg := NewConfigFromYaml(yamlFromFile(t))

	all := cfg.OutputAll()

	require.True(t, len(all) > 0, "results did not contain any cm's")

	result, exists := all["int"]
	require.True(t, exists, "default did not exist in results")
	require.Equal(t, "foo", result.Name, "name did not match")
	require.Equal(t, "system", result.Namespace, "namespace did not fatch")
	require.Equal(t, "an-override", result.Data["setting1"], "setting1 did not match")
	require.Equal(t, "value2", result.Data["setting2"], "setting1 did not match")

}
