//go:generate go run generate_help.go

package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	clis := []struct {
		cmd string
		out string
	}{
		{"./cmd/vbantxt/", "vbantxt-help.md"},
		{"./cmd/vbantxtctl/", "vbantxtctl-help.md"},
	}

	for _, cli := range clis {
		cmdArgs := []string{"run", cli.cmd, "--help"}
		helpCmd := exec.Command("go", cmdArgs...)
		helpCmd.Dir = ".."
		out, err := helpCmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		// Wrap output in markdown console code block
		wrapped := append([]byte("```console\n"), out...)
		wrapped = append(wrapped, []byte("\n```\n")...)

		if err := os.WriteFile(cli.out, wrapped, 0o644); err != nil {
			log.Fatal(err)
		}
	}
}
