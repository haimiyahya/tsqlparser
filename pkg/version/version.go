// Package version provides version information for tsqlparser.
//
// The version is kept in sync with the VERSION file at the repository root.
// When releasing, update the VERSION file and regenerate this file using:
//
//	go generate ./version
//
// Or simply update the Version constant manually.
package version

import (
	_ "embed"
	"strings"
)

//go:embed version.txt
var versionFile string

// Version is the current version of tsqlparser.
// This is embedded from version.txt at compile time.
var Version = strings.TrimSpace(versionFile)

// String returns the version string.
func String() string {
	return Version
}

// Full returns a full version string with the package name.
func Full() string {
	return "tsqlparser version " + Version
}
