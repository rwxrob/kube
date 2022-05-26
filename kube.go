package kube

import (
	"os"
	"path/filepath"

	"github.com/rwxrob/fs/file"
	"gopkg.in/yaml.v3"
)

// DefConfFile returns the full path to the default config file
// depending on the OS. Returns empty if nothing found.
func DefConfFile() string {
	d, err := os.UserHomeDir()

	if err != nil {
		return ""
	}
	path := filepath.Join(d, `.kube`, `config`)
	if file.Exists(path) {
		return path
	}
	return ""
}

type Config struct {
	O map[string]any `yaml:",inline"`
}

func (c *Config) load(path string) error {
	if !file.Exists(path) {
		return nil
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &c.O)
	if err != nil {
		return err
	}

	return nil
}

// Load follows the same conventions as kubectl for loading Kubernetes
// configuration YAML data, first from all files in any KUBECONFIG
// environment variable path are loaded with the values of the last file
// having priority over those before it. If KUBECONFIG is not set or
// empty then .kube/config file in the users' home directory is loaded
// instead. Any files that do not exist are silently ignored.
func (c *Config) Load() error {
	kubeconfig := os.Getenv(`KUBECONFIG`)

	if kubeconfig == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		kubeconfig = filepath.Join(home, `.kube`, `config`)
	}

	paths := filepath.SplitList(kubeconfig)
	for _, path := range paths {
		c.load(path)
	}

	return nil
}

// CurrentContext first loads the kubeconfig (see Config.Load) and then
// returns the current-context value or empty string if not found.
func CurrentContext() string {
	c := Config{}
	c.Load()
	if a, has := c.O[`current-context`]; has {
		return a.(string)
	}
	return ""
}
