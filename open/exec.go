// +build !windows,!darwin

package open

import (
	"bytes"
	"os/exec"
	"strings"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

func open(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}

func openWait(input string) *exec.Cmd {
	buf := &bytes.Buffer{}
	cmd := exec.Command("xdg-mime query default text/plain")
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return err
	}
	path := strings.TrimSpace(buf.String())
	return exec.Command(path, input)
}

func openWithWait(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}
