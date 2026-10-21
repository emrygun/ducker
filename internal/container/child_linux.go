//go:build linux

package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/emreuygun/ducker/internal/cgroups"
)

// Child runs inside the new namespaces created by Run. It sets the
// hostname, chroots into the rootfs, mounts /proc, applies cgroup
// limits, then execs the requested command.
func Child(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: ducker child <rootfs> <command> [args...]")
	}

	rootfs := args[0]
	command := args[1]
	commandArgs := args[2:]

	if err := syscall.Sethostname([]byte("ducker")); err != nil {
		return fmt.Errorf("sethostname: %w", err)
	}

	if err := cgroups.Apply(os.Getpid()); err != nil {
		return fmt.Errorf("cgroups: %w", err)
	}

	if err := syscall.Chroot(rootfs); err != nil {
		return fmt.Errorf("chroot: %w", err)
	}
	if err := os.Chdir("/"); err != nil {
		return fmt.Errorf("chdir: %w", err)
	}

	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return fmt.Errorf("mount proc: %w", err)
	}
	defer syscall.Unmount("/proc", 0)

	cmd := exec.Command(command, commandArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
