package configurator

import (
	"github.com/ghodss/yaml"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Config represents a default ConfigMap
// and any applicable overrides
type Config struct {
	Name        string                       `yaml:"name"`
	Namespace   string                       `yaml:"namespace"`
	Default     map[string]string            `yaml:"default"`
	Overrides   map[string]map[string]string `yaml:"overrides"`
	Annotations map[string]string            `yaml:"annotations,omitempty"`
}

// NewConfigFromYaml reads yaml from a byte slice
// and returns a populated config.
func NewConfigFromYaml(in []byte) Config {
	var cfg Config

	yaml.Unmarshal(in, &cfg)

	return cfg
}

// OutputAll returns a Default ConfigMap and
// any merged overrides
func (c Config) OutputAll() map[string]v1.ConfigMap {
	results := make(map[string]v1.ConfigMap)

	base := v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.Name,
			Namespace:   c.Namespace,
			Annotations: c.Annotations,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: c.Default,
	}
	results["default"] = base

	for override, merge := range c.Overrides {
		o := v1.ConfigMap{}
		base.DeepCopyInto(&o)

		for k, v := range merge {
			o.Data[k] = v
		}

		results[override] = o
	}
	return results
}
