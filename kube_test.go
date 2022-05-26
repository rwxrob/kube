package kube_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwxrob/kube"
)

/*
func ExampleDefConfFile() {
	fmt.Println(kube.DefConfFile())
	// Output: /home/rwxrob/.kube/config
}
*/

/*
func ExampleCurrentContext() {
	fmt.Println(kube.CurrentContext())
	// Output: kind-kind
}
*/

func ExampleConfig_Load_kubeconfig() {

	path := []string{
		filepath.Join(`testdata`, `conf1.yaml`),
		filepath.Join(`testdata`, `conf2.yaml`),
	}
	pathstr := strings.Join(path, string(filepath.ListSeparator))
	log.Println(pathstr)
	orig := os.Getenv(`KUBECONFIG`)
	os.Setenv(`KUBECONFIG`, pathstr)
	defer func() {
		if orig != "" {
			os.Setenv(`KUBECONFIG`, orig)
		}
	}()

	c := kube.Config{}
	c.Load()
	fmt.Println(c.O[`current-context`])

	// Output:
	// foo
}

func ExampleCurrentContext() {

	path := []string{
		filepath.Join(`testdata`, `conf1.yaml`),
		filepath.Join(`testdata`, `conf2.yaml`),
	}
	pathstr := strings.Join(path, string(filepath.ListSeparator))
	log.Println(pathstr)
	orig := os.Getenv(`KUBECONFIG`)
	os.Setenv(`KUBECONFIG`, pathstr)
	defer func() {
		if orig != "" {
			os.Setenv(`KUBECONFIG`, orig)
		}
	}()

	fmt.Println(kube.CurrentContext())

	// Output:
	// foo
}
