package hooks

import (
	"os"
	"regexp"

	"github.com/tftp-go-team/libgotftp/src"
)

var pathEscape = regexp.MustCompile("\\.\\.\\/")

var FileHook = HookComponents{
	func(path string, _ tftp.Request) (*HookResult, error) {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		// get the file size
		stat, err := file.Stat()
		if err != nil {
			return nil, err
		}
		return &HookResult{
			Stdout: file,
			Stderr: nil,
			Length: int(stat.Size()),
		}, nil
	},
	func(s string) string {
		return pathEscape.ReplaceAllLiteralString(s, "")
	},
}
