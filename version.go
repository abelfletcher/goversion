package goversion

import (
	"strconv"
	"strings"
)

// Version is stored in the application's VERSION var
type Version struct {
	major VersionComponent
	minor VersionComponent
	patch VersionComponent
	rc VersionComponent
	beta VersionComponent

	hasRc bool
	hasBeta bool
	serialized string
}

// VERSION sets the version for the application
func VERSION(version string) *Version {
	return deserialize(version)
}

// Major returns the major version component
func(v *Version) Major() VersionComponent {
	return v.major
}

// MajorIs compares the major version to a known version
func(v *Version) MajorIs(what int) bool {
	if VersionComponent(what) == v.major {
		return true
	}

	return false
}

// Minor returns the minor version component
func(v *Version) Minor() VersionComponent {
	return v.minor
}

// MinorIs compares the minor version to a known version
func(v *Version) MinorIs(what int) bool {
	if VersionComponent(what) == v.minor {
		return true
	}

	return false
}

// Patch returns the patch version component
func(v *Version) Patch() VersionComponent {
	return v.patch
}

// PatchIs compares the patch versiont to a known version
func(v *Version) PatchIs(what int) bool {
	if VersionComponent(what) == v.patch {
		return true
	}

	return false
}

// IsRc returns whether or not this is an rc version
func(v *Version) IsRc() bool {
	return v.hasRc
}

// Rc returns the release candidate component
func(v *Version) Rc() VersionComponent {
	return v.rc
}

// RcIs compares the rc version to a known version
func(v *Version) RcIs(what int) bool {
	if v.hasRc {
		if VersionComponent(what) == v.rc {
			return true
		}
	}

	return false
}

// IsBeta returns whether or not this is a beta version
func(v *Version) IsBeta() bool {
	return v.hasBeta
}

// Beta returns the beta component
func(v *Version) Beta() VersionComponent {
	return v.beta
}

// BetaIs compares the beta version to a known version
func(v *Version) BetaIs(what int) bool {
	if v.hasBeta {
		if VersionComponent(what) == v.beta {
			return true
		}
	}

	return false
}

// String returns the full version string
func(v *Version) String() string {
	if len(v.serialized) == 0 {
		v.serialized = serialize(v)
	}

	return v.serialized
}

// serialize returns a string serialization
func serialize(v *Version) string {
	s := v.Major().String()+"."+v.Minor().String()+
		"."+v.Patch().String()

	if v.hasRc {
		s += "-rc"+v.Rc().String()
	}

	if v.hasBeta {
		s += "b"+v.Beta().String()
	}

	return s
}

// deserialize returns a Version from string
func deserialize(vs string) *Version {
	var idx int
	var tmp uint64
	var err error

	index := func(delim string) {
		idx = strings.Index(vs, delim)
	}

	trim := func(delim string) {
		vs = vs[idx+len(delim):]
	}

	parse := func() {
		tmp, err = strconv.ParseUint(vs[:idx], 10, 64)
	}

	contains := func(what string) bool {
		return strings.Contains(vs, what)
	}

	set := func(what *VersionComponent, delim string) {
		if delim != "" {
			index(delim)
			parse()
			trim(delim)
		} else {
			idx = len(vs)
			parse()
		}

		if err != nil {
			*what = VersionComponent(0)
		} else {
			*what = VersionComponent(tmp)
		}

	}

	v := new(Version)

	if contains("-rc") {
		v.hasRc = true
	}

	if contains("b") {
		v.hasBeta = true
	}

	if contains(".") {
		set(&v.major, ".")

		if contains(".") {
			set(&v.minor, ".")

			if v.hasRc {
				set(&v.patch, "-rc")

				if v.hasBeta {
					set(&v.rc, "b")
					set(&v.beta, "")
				} else {
					set(&v.rc, "")
				}
			} else {
				if v.hasBeta {
					set(&v.patch, "b")
					set(&v.beta, "")
				} else {
					set(&v.patch, "")
				}
			}
		} else {
			v.minor = 0
			v.patch = 0
			v.hasRc = false
			v.hasBeta = false
		}
	} else {
		v.major = 0
		v.minor = 0
		v.patch = 0
		v.hasRc = false
		v.hasBeta = false
	}

	return v
}

// VersionComponent tracks numerical values
type VersionComponent uint64

// Return the component as string
func(v VersionComponent) String() string {
	return strconv.FormatUint(v.Uint64(), 10)
}

// Return the component as uint8
func(v VersionComponent) Uint8() uint8 {
	return uint8(v)
}

// Return the component as uint16
func(v VersionComponent) Uint16() uint16 {
	return uint16(v)
}

// Return the component as uint32
func(v VersionComponent) Uint32() uint32 {
	return uint32(v)
}

// Return the component as uint64
func(v VersionComponent) Uint64() uint64 {
	return uint64(v)
}

// Return the component as int8
func(v VersionComponent) Int8() int8 {
	return int8(v)
}

// Return the component as int16
func(v VersionComponent) Int16() int16 {
	return int16(v)
}

// Return the component as int32
func(v VersionComponent) Int32() int32 {
	return int32(v)
}

// Return the component as int64
func(v VersionComponent) Int64() int64 {
	return int64(v)
}

// Return the component as int
func(v VersionComponent) Int() int {
	return int(v)
}
