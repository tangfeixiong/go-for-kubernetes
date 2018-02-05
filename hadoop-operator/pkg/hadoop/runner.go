/*
   Inspired by:
    https://github.com/helm/helm-classic/blob/master/kubectl/kubectl.go
*/
// package kubectl
package hadoop

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Path is the path of the kubectl binary
var Path = "sh"

// Runner is an interface to wrap kubectl convenience methods
type Runner interface {
	//	// ClusterInfo returns Kubernetes cluster info
	//	ClusterInfo() ([]byte, error)
	//	// Create uploads a chart to Kubernetes
	//	Create([]byte, string) ([]byte, error)
	//	// Delete removes a chart from Kubernetes.
	//	Delete(string, string, string) ([]byte, error)
	//	// Get returns Kubernetes resources
	//	Get([]byte, string) ([]byte, error)
	Format(string, string, string) error
}

// RealRunner implements Runner to execute kubectl commands
type RealRunner struct{}

// PrintRunner implements Runner to return a []byte of the command to be executed
type PrintRunner struct{}

// Client stores the instance of Runner
var Client Runner = RealRunner{}

func commandToString(cmd *exec.Cmd) string {
	var stdin string

	if cmd.Stdin != nil {
		b, _ := ioutil.ReadAll(cmd.Stdin)
		stdin = fmt.Sprintf("< %s", string(b))
	}

	return fmt.Sprintf("[CMD] %s %s", strings.Join(cmd.Args, " "), stdin)
}
