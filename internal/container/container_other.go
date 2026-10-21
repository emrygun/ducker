//go:build !linux

package container

import "fmt"

// Run and Child rely on Linux-only namespace/cgroup syscalls, so ducker
// only works on Linux. These stubs let the project build (e.g. for
// editing on macOS) but fail clearly at runtime on other platforms.

func Run(args []string) error {
	return fmt.Errorf("ducker only supports linux (run it in a Linux VM or container)")
}

func Child(args []string) error {
	return fmt.Errorf("ducker only supports linux (run it in a Linux VM or container)")
}
