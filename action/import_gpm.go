package action

import (
	"github.com/4396/glide/gpm"
	"github.com/4396/glide/msg"
)

// ImportGPM imports a GPM file.
func ImportGPM(dest string) {
	base := "."
	config := EnsureConfig()
	if !gpm.Has(base) {
		msg.Die("No GPM Godeps file found.")
	}
	deps, err := gpm.Parse(base)
	if err != nil {
		msg.Die("Failed to extract GPM Godeps file: %s", err)
	}
	appendImports(deps, config)
	writeConfigToFileOrStdout(config, dest)
}
