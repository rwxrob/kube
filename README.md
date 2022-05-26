# 🌳 Kubernetes Helper Bonzai Command Branch

[![GoDoc](https://godoc.org/github.com/rwxrob/kube?status.svg)](https://godoc.org/github.com/rwxrob/kube)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Install

This command can be installed as a standalone program or composed into a
Bonzai™ command tree.

Standalone

```
go install github.com/rwxrob/kube/cmd/kube@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/kube/z"
	"github.com/rwxrob/kube"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, kube.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C kube kube
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
