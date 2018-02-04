package configurator

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGenerateDefault(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	input := yamlFromFile(t)
	err := Generate(input, "default", buf)
	require.NoError(t, err)
	yaml := buf.String()
	require.Contains(t, yaml, "kind: ConfigMap")
	require.Contains(t, yaml, "setting1: value1")
}

func TestCanGenerateOverride(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	input := yamlFromFile(t)
	err := Generate(input, "int", buf)
	require.NoError(t, err)
	yaml := buf.String()
	require.Contains(t, yaml, "kind: ConfigMap")
	require.Contains(t, yaml, "setting1: an-override")
}

func TestGenerateErrorsOnInvalidEnv(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	input := yamlFromFile(t)
	err := Generate(input, "foobar", buf)
	require.Error(t, err, "no error returned for invalid env")
}
