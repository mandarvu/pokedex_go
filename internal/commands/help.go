package commands

import "fmt"

func HelpText(supportedCommands map[string]Command) string {
	hstr := []byte{}

	for k, v := range supportedCommands {
		s := fmt.Sprintf("%s: %s\n", k, v.description)
		hstr = fmt.Append(hstr, s)
	}

	return string(hstr)
}
