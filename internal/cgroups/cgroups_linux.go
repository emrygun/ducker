//go:build linux

// Package cgroups applies basic resource limits (memory, pids) to a
// container process via cgroups v1. It's intentionally minimal: a toy
// implementation, not production-ready isolation.
package cgroups

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	cgroupRoot = "/sys/fs/cgroup"
	memLimit   = 100 * 1024 * 1024 // 100MB
	pidsLimit  = 64
)

// Apply creates a "ducker" cgroup under the memory and pids controllers
// and adds pid to it.
func Apply(pid int) error {
	if err := applyController("memory", "memory.limit_in_bytes", memLimit, pid); err != nil {
		return err
	}
	if err := applyController("pids", "pids.max", pidsLimit, pid); err != nil {
		return err
	}
	return nil
}

func applyController(controller, limitFile string, limit, pid int) error {
	dir := filepath.Join(cgroupRoot, controller, "ducker")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}

	if err := os.WriteFile(filepath.Join(dir, limitFile), []byte(strconv.Itoa(limit)), 0644); err != nil {
		return fmt.Errorf("write %s: %w", limitFile, err)
	}

	if err := os.WriteFile(filepath.Join(dir, "cgroup.procs"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("write cgroup.procs: %w", err)
	}

	return nil
}
