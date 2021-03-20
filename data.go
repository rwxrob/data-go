package data

import (
	"os"
	"os/user"
	"path/filepath"
)

// FileName is the default file name to use. By default this is
// "data.yml" and will be joined with the appropriate suffix for the
// structured data type.
var FileName string = `data.yml`

// ForceDir forces Dir() to return its value when set. This is mostly
// just there for mock testing since Dir() depends (in part) on the
// current user home directory. Also see cmdtab.User.
var ForceDir string

// Dir returns the main data directory into which the specific data
// store directory and file will be created. Dir is always
// XDG_DATA_HOME if detected or ~/.local/share if not.  During
// testing this can, however, be forced to be something else by setting
// ForceDir explicitly (which will always override everything else).
func Dir() string {
	if ForceDir != "" {
		return ForceDir
	}
	dir := os.Getenv("XDG_DATA_HOME")
	if dir != "" {
		return dir
	}
	usr, _ := user.Current()
	if usr != nil {
		return filepath.Join(usr.HomeDir, ".local", "share")
	}
	return ""
}

// Path returns the full path to the Dir() combined with the arguments.
func Path(args ...string) string {
	p := []string{Dir()}
	p = append(p, args...)
	return filepath.Join(p...)
}
