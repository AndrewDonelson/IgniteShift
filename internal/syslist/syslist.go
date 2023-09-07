// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syslist // changed from go -> build

import (
	"errors"
	"strings"
) // Note that this file is read by internal/goarch/gengoarch.go and by
// internal/goos/gengoos.go. If you change this file, look at those
// files as well.

// knownOS is the list of past, present, and future known GOOS values.
// Do not remove from this list, as it is used for filename matching.
// If you add an entry to this list, look at unixOS, below.
var KnownOS = map[string]bool{
	"aix":       true,
	"android":   true,
	"darwin":    true,
	"dragonfly": true,
	"freebsd":   true,
	"hurd":      true,
	"illumos":   true,
	"ios":       true,
	"js":        true,
	"linux":     true,
	"nacl":      true,
	"netbsd":    true,
	"openbsd":   true,
	"plan9":     true,
	"solaris":   true,
	"windows":   true,
	"zos":       true,
}

// unixOS is the set of GOOS values matched by the "unix" build tag.
// This is not used for filename matching.
// This list also appears in cmd/dist/build.go.
var UnixOS = map[string]bool{
	"aix":       true,
	"android":   true,
	"darwin":    true,
	"dragonfly": true,
	"freebsd":   true,
	"hurd":      true,
	"illumos":   true,
	"ios":       true,
	"linux":     true,
	"netbsd":    true,
	"openbsd":   true,
	"solaris":   true,
}

// knownArch is the list of past, present, and future known GOARCH values.
// Do not remove from this list, as it is used for filename matching.
var KnownArch = map[string]bool{
	"386":         true,
	"amd64":       true,
	"amd64p32":    true,
	"arm":         true,
	"armbe":       true,
	"arm64":       true,
	"arm64be":     true,
	"loong64":     true,
	"mips":        true,
	"mipsle":      true,
	"mips64":      true,
	"mips64le":    true,
	"mips64p32":   true,
	"mips64p32le": true,
	"ppc":         true,
	"ppc64":       true,
	"ppc64le":     true,
	"riscv":       true,
	"riscv64":     true,
	"s390":        true,
	"s390x":       true,
	"sparc":       true,
	"sparc64":     true,
	"wasm":        true,
}

// Operating systems enumeration
type OS int

const (
	OS_UNKNOWN OS = iota
	OS_LINUX
	OS_DARWIN
	OS_WINDOWS
)

// String returns the string representation of the OS
func (os OS) String() string {
	switch os {
	case OS_LINUX:
		return "linux"
	case OS_DARWIN:
		return "darwin"
	case OS_WINDOWS:
		return "windows"
	default:
		return "unknown"
	}
}

// Value returns the integer representation of the OS
func (os OS) Value() int {
	return int(os)
}

// OSFromString returns the OS from a string
func OSFromString(osStr string) (OS, error) {
	switch strings.ToLower(osStr) {
	case "linux":
		return OS_LINUX, nil
	case "darwin":
		return OS_DARWIN, nil
	case "windows":
		return OS_WINDOWS, nil
	default:
		return OS_UNKNOWN, errors.New("invalid OS string")
	}
}

// OSFromValue returns the OS from an integer
func OSFromValue(val int) (OS, error) {
	if val >= 0 && val <= int(OS_WINDOWS) {
		return OS(val), nil
	}
	return OS_UNKNOWN, errors.New("invalid OS value")
}

// Architecture enumeration
type Arch int

const (
	ARCH_UNKNOWN Arch = iota
	ARCH_AMD64
	ARCH_ARM64
)

// String returns the string representation of the architecture
func (arch Arch) String() string {
	switch arch {
	case ARCH_AMD64:
		return "amd64"
	case ARCH_ARM64:
		return "arm64"
	default:
		return "unknown"
	}
}

// Value returns the integer representation of the architecture
func (arch Arch) Value() int {
	return int(arch)
}

// ArchFromString returns the architecture from a string
func ArchFromString(archStr string) (Arch, error) {
	switch strings.ToLower(archStr) {
	case "amd64":
		return ARCH_AMD64, nil
	case "arm64":
		return ARCH_ARM64, nil
	default:
		return ARCH_UNKNOWN, errors.New("invalid Arch string")
	}
}

// ArchFromValue returns the architecture from an integer
func ArchFromValue(val int) (Arch, error) {
	if val >= 0 && val <= int(ARCH_ARM64) {
		return Arch(val), nil
	}
	return ARCH_UNKNOWN, errors.New("invalid Arch value")
}

// Validate checks if the osarch string is a valid combination according to the defined OS and Arch enumerations
func Validate(osarch string) bool {
	parts := strings.Split(osarch, "/")
	if len(parts) != 2 {
		return false
	}

	osStr, archStr := parts[0], parts[1]

	var validOS, validArch bool

	switch osStr {
	case OS_LINUX.String(), OS_DARWIN.String(), OS_WINDOWS.String():
		validOS = true
	}

	switch archStr {
	case ARCH_AMD64.String(), ARCH_ARM64.String():
		validArch = true
	}

	return validOS && validArch
}
