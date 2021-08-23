package cmd

import (
	"embed"
	"log"
)

func Execute(r embed.FS) {
	root = r
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
