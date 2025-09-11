//go:build !windows

package hiddify

import (
	"os/exec"
	"syscall"
)

func setpgid(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}
