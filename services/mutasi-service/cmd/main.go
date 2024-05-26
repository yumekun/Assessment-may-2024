package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Usage = help
	flag.Parse()

	cmds := map[string]func(){
		"help":  help,
		"start": start,
	}

	if cmdFunc, ok := cmds[flag.Arg(0)]; ok {
		cmdFunc()
	} else {
		help()
		os.Exit(2)
	}
}

func help() {
	divider := "| %s | %s |\n"
	header := "| %-20s | %-40s |\n"
	row := "| %-20s | %-40s |\n"

	output :=
		fmt.Sprintf(header, "Usage", "Description") +
			fmt.Sprintf(divider, strings.Repeat("-", 20), strings.Repeat("-", 40)) +
			fmt.Sprintf(row, "`journal-service` help", "show this message") +
			fmt.Sprintf(row, "`journal-service` start", "start the server") +
			fmt.Sprintf(divider, strings.Repeat("_", 20), strings.Repeat("_", 40))

	fmt.Fprintln(os.Stderr, output)
}
