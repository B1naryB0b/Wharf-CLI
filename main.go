package main

import (
	"Wharf-CLI/cmd"
)

var Version = "dev"

func main() {
	cmd.Version = Version
	cmd.Execute()
}
