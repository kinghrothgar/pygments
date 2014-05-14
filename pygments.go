// Pygments wrapper for golang. Pygments is a syntax highlighter

package pygments

import (
	"bytes"
	"os/exec"
	"strings"
	"errors"
)

type Options map[string]string

var (
	bin = "pygmentize"
)

func Binary(path string) {
	bin = path
}

func Which() string {
	return bin
}

func Highlight(code string, lexer string, format string, options Options) (string, error) {

	if _, err := exec.LookPath(bin); err != nil {
		return "", errors.New("Could not find '"+bin+"'")
	}

	optionsString := ""
	for name, value := range options  {
		optionsString += name
		if value != "" {
			optionsString += "=" + value
		}
	}
	strings.TrimSuffix(optionsString, ",")

	var cmd *exec.Cmd
	if len(optionsString) > 0 {
		cmd = exec.Command(bin, "-l"+lexer, "-f"+format, "-O "+optionsString)
	} else {
		cmd = exec.Command(bin, "-l"+lexer, "-f"+format)
	}
	cmd.Stdin = strings.NewReader(code)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", errors.New("Failed to run highlight command with error: "+err.Error())
	}

	return out.String(), nil
}
