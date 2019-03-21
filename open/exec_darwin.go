// +build darwin

package open

import (
	"os/exec"
)

func open(input string) *exec.Cmd {
	return exec.Command("open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command("open", "-a", appName, input)
}

func openWait(input string) *exec.Cmd {
	return exec.Command("open", "-W", "-n", input)
}

func openWithWait(input string, appName string) *exec.Cmd {
	return exec.Command("open", "-W", "-n", "-a", appName, input)
}
