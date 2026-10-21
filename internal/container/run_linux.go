//go:build linux

package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Run is the entrypoint for `ducker run <rootfs> <cmd> [args...]`.
// It re-execs the ducker binary as `child` inside a fresh set of namespaces.
func Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: ducker run <rootfs> <command> [args...]")
	}

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	return cmd.Run()
}
