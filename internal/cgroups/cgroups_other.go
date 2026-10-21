//go:build !linux

package cgroups

import "fmt"

func Apply(pid int) error {
	return fmt.Errorf("cgroups only supported on linux")
}
