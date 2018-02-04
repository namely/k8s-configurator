package configurator

import (
	"fmt"
	"io"

	"github.com/ghodss/yaml"
)

// Generate generates a ConfigMap for the given environment to the
// specified output. Use 'default' for the default ConfigMap.
func Generate(input []byte, env string, output io.Writer) error {

	cfg := NewConfigFromYaml(input)

	all := cfg.OutputAll()

	e, exists := all[env]
	if !exists {
		return fmt.Errorf("requested env %v does not exist in input yaml", env)
	}

	d, err := yaml.Marshal(e)
	if err != nil {
		return err
	}
	fmt.Fprintf(output, "%s", d)

	return nil
}
