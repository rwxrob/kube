package kube

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

func init() {
	Z.Dynamic[`kubecontext`] = CurrentContext
}

var Cmd = &Z.Cmd{
	Name:      `kube`,
	Summary:   `({{.Version}}) Kubernetes utility commands`,
	Version:   `v0.1.0`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:rwxrob/kube.git`,
	Issues:    `github.com/rwxrob/kube/issues`,

	Commands: []*Z.Cmd{help.Cmd, contextCmd},
	MinArgs:  1,

	Description: `
		The {{cmd .Name}} commands provide user-centric utilities and tools
		for working with Kubernetes clusters and applications.
	`,
}

var contextCmd = &Z.Cmd{
	Name:     `context`,
	Summary:  `print current Kubernetes context`,
	Commands: []*Z.Cmd{help.Cmd},
	Source:   `https://github.com/rwxrob/kube`,
	Issues:   `https://github.com/rwxrob/kube/issues`,
	Description: `
		The {{cmd .Name}} commands prints the current kubeconfig
		current-context value (currently: {{ kubecontext }}).`,
	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(CurrentContext)
		return nil
	},
}
