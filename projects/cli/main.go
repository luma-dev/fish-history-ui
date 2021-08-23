package main

import (
	"embed"

	"github.com/luma-dev/fish-history-ui/projects/cli/cmd"
)

//go:embed web-ui-dist/* web-ui-dist/**/*
var root embed.FS

func main() {
	cmd.Execute(root)
}
