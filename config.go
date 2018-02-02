package configurator

import (
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

// Config represents a default ConfigMap
// and any applicable overrides
type Config struct {
	Name      string
	Namespace string
	Default   map[string]string
	Overrides map[string]map[string]string
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
func (c Config) OutputAll() map[string]core.ConfigMap {
	results := make(map[string]core.ConfigMap)

	base := core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:              c.Name,
			Namespace:         c.Namespace,
			CreationTimestamp: metav1.Time{},
			DeletionTimestamp: &metav1.Time{},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: c.Default,
	}
	results["default"] = base

	for override, merge := range c.Overrides {
		o := core.ConfigMap{}
		base.DeepCopyInto(&o)

		for k, v := range merge {
			o.Data[k] = v
		}

		results[override] = o
	}
	return results
}
