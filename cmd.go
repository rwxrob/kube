package kube

import (
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `kube`,
	Summary:  `({{.Version}}) Kubernetes utility commands`,
	Commands: []*Z.Cmd{help.Cmd},
	MinArgs:  1,
	Version:  `v0.0.1`,
	Source:   `https://github.com/rwxrob/kube`,
	Issues:   `https://github.com/rwxrob/kube/issues`,
	Description: `
		The {{cmd .Name}} commands provide user-centric utilities and tools
		for working with Kubernetes clusters and applications.
	`,
}
