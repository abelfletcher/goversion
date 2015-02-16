package goversion

import (
	"strconv"
	"strings"
)

// Version is stored in the application's VERSION var
type Version struct {
	major Major
	minor Minor
	patch Patch
	rc    Rc
	beta  Beta

	hasRc      bool
	hasBeta    bool
	serialized string
}

// VERSION sets the version for the application
func VERSION(version string) *Version {
	return deserialize(version)
}

func NewVersion(version interface{}) *Version {
	switch version.(type) {
	case string:
		return VERSION(version.(string))
	case Version:
		v := version.(Version)
		return NewVersion(&v)
	case *Version:
		return NewVersion(version.(*Version).String())
	}

	return NewVersion("")
}

// Is compares the current version to a known version
func (v *Version) Is(version interface{}) bool {
	v1 := NewVersion(version)
	if !v.major.is(v1.major) {
		return false
	}

	if !v.minor.is(v1.minor) {
		return false
	}

	if !v.patch.is(v1.patch) {
		return false
	}

	if v.hasRc != v1.hasRc {
		return false
	}

	if !v.rc.is(v1.rc) {
		return false
	}

	if v.hasBeta != v1.hasBeta {
		return false
	}

	if !v.beta.is(v1.beta) {
		return false
	}

	return true
}

// Equals aliases Is()
func (v *Version) Equals(version interface{}) bool {
	return v.Is(version)
}

// LessThan compares current version to known version
func (v *Version) LessThan(version interface{}) bool {
	v1 := NewVersion(version)

	if v.major.less(v1.major) {
		return true
	}

	if v.major.is(v1.major) {
		if v.minor.less(v1.minor) {
			return true
		}

		if v.minor.is(v1.minor) {
			if v.patch.less(v1.patch) {
				return true
			}

			if v.patch.is(v1.patch) {
				if v.hasRc {
					if v1.hasRc {
						if v.rc.less(v1.rc) {
							return true
						}

						if v.rc.is(v1.rc) {
							if v.hasBeta {
								if v1.hasBeta {
									if v.beta.less(v1.beta) {
										return true
									}
								}
							} else {
								if v1.hasBeta {
									return true
								}
							}
						}
					}
				} else {
					if v1.hasRc {
						return true
					}

					if v.hasBeta {
						if v1.hasBeta {
							if v.beta.less(v1.beta) {
								return true
							}
						}
					} else {
						if v1.hasBeta {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

// Lt aliases LessThan
func (v *Version) Lt(version interface{}) bool {
	return v.LessThan(version)
}

// LessThanOrEqualTo compares current version to known version
func (v *Version) LessThanOrEqualTo(version interface{}) bool {
	if v.Is(version) || v.LessThan(version) {
		return true
	}

	return false
}

// Lte aliases LessThanOrEqualTo
func (v *Version) Lte(version interface{}) bool {
	return v.LessThanOrEqualTo(version)
}

// GreaterThan compares current version to known version
func (v *Version) GreaterThan(version interface{}) bool {
	v1 := NewVersion(version)

	if v.major.greater(v1.major) {
		return true
	}

	if v.major.is(v1.major) {
		if v.minor.greater(v1.minor) {
			return true
		}

		if v.minor.is(v1.minor) {
			if v.patch.greater(v1.patch) {
				return true
			}

			if v.patch.is(v1.patch) {
				if v.hasRc {
					if v1.hasRc {
						if v.rc.greater(v1.rc) {
							return true
						}

						if v.rc.is(v1.rc) {
							if v.hasBeta {
								if v1.hasBeta {
									if v.beta.greater(v1.beta) {
										return true
									}
								} else {
									return true
								}
							}
						}
					} else {
						return true
					}
				} else {
					if !v1.hasRc {
						if v.hasBeta {
							if v1.hasBeta {
								if v.beta.greater(v1.beta) {
									return true
								}
							} else {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

// Gt aliases GreaterThan
func (v *Version) Gt(version interface{}) bool {
	return v.GreaterThan(version)
}

// GreaterThanOrEqualTo compares current to known version
func (v *Version) GreaterThanOrEqualTo(version interface{}) bool {
	if v.Is(version) || v.GreaterThan(version) {
		return true
	}

	return false
}

// Gte aliases GreaterThanOrEqualTo
func (v *Version) Gte(version interface{}) bool {
	return v.GreaterThanOrEqualTo(version)
}

// Major returns the major version component
func (v *Version) Major() Major {
	return v.major
}

// MajorIs compares the major version to a known version
func (v *Version) MajorIs(what int) bool {
	if VersionComponent(what) == v.major.VersionComponent {
		return true
	}

	return false
}

// Minor returns the minor version component
func (v *Version) Minor() Minor {
	return v.minor
}

// MinorIs compares the minor version to a known version
func (v *Version) MinorIs(what int) bool {
	if VersionComponent(what) == v.minor.VersionComponent {
		return true
	}

	return false
}

// Patch returns the patch version component
func (v *Version) Patch() Patch {
	return v.patch
}

// PatchIs compares the patch versiont to a known version
func (v *Version) PatchIs(what int) bool {
	if VersionComponent(what) == v.patch.VersionComponent {
		return true
	}

	return false
}

// IsRc returns whether or not this is an rc version
func (v *Version) IsRc() bool {
	return v.hasRc
}

// Rc returns the release candidate component
func (v *Version) Rc() Rc {
	return v.rc
}

// RcIs compares the rc version to a known version
func (v *Version) RcIs(what int) bool {
	if v.hasRc {
		if VersionComponent(what) == v.rc.VersionComponent {
			return true
		}
	}

	return false
}

// IsBeta returns whether or not this is a beta version
func (v *Version) IsBeta() bool {
	return v.hasBeta
}

// Beta returns the beta component
func (v *Version) Beta() Beta {
	return v.beta
}

// BetaIs compares the beta version to a known version
func (v *Version) BetaIs(what int) bool {
	if v.hasBeta {
		if VersionComponent(what) == v.beta.VersionComponent {
			return true
		}
	}

	return false
}

// String returns the full version string
func (v *Version) String() string {
	if len(v.serialized) == 0 {
		v.serialized = serialize(v)
	}

	return v.serialized
}

// serialize returns a string serialization
func serialize(v *Version) string {
	s := v.Major().String() + "." + v.Minor().String() +
		"." + v.Patch().String()

	if v.hasRc {
		s += "-rc" + v.Rc().String()
	}

	if v.hasBeta {
		s += "b" + v.Beta().String()
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
		set(&v.major.VersionComponent, ".")

		if contains(".") {
			set(&v.minor.VersionComponent, ".")

			if v.hasRc {
				set(&v.patch.VersionComponent, "-rc")

				if v.hasBeta {
					set(&v.rc.VersionComponent, "b")
					set(&v.beta.VersionComponent, "")
				} else {
					set(&v.rc.VersionComponent, "")
				}
			} else {
				if v.hasBeta {
					set(&v.patch.VersionComponent, "b")
					set(&v.beta.VersionComponent, "")
				} else {
					set(&v.patch.VersionComponent, "")
				}
			}
		} else {
			v.minor.VersionComponent = 0
			v.patch.VersionComponent = 0
			v.hasRc = false
			v.hasBeta = false
		}
	} else {
		v.major.VersionComponent = 0
		v.minor.VersionComponent = 0
		v.patch.VersionComponent = 0
		v.hasRc = false
		v.hasBeta = false
	}

	return v
}

// VersionComponent tracks numerical values
type VersionComponent uint64
type Major struct{ VersionComponent }
type Minor struct{ VersionComponent }
type Patch struct{ VersionComponent }
type Rc struct{ VersionComponent }
type Beta struct{ VersionComponent }

type vComponent interface {
	Component() VersionComponent
}

// Returns self
func (v VersionComponent) Component() VersionComponent {
	return v
}

// Return the component as string
func (v VersionComponent) String() string {
	return strconv.FormatUint(v.Uint64(), 10)
}

// Return the component as uint8
func (v VersionComponent) Uint8() uint8 {
	return uint8(v)
}

// Return the component as uint16
func (v VersionComponent) Uint16() uint16 {
	return uint16(v)
}

// Return the component as uint32
func (v VersionComponent) Uint32() uint32 {
	return uint32(v)
}

// Return the component as uint64
func (v VersionComponent) Uint64() uint64 {
	return uint64(v)
}

// Return the component as int8
func (v VersionComponent) Int8() int8 {
	return int8(v)
}

// Return the component as int16
func (v VersionComponent) Int16() int16 {
	return int16(v)
}

// Return the component as int32
func (v VersionComponent) Int32() int32 {
	return int32(v)
}

// Return the component as int64
func (v VersionComponent) Int64() int64 {
	return int64(v)
}

// Return the component as int
func (v VersionComponent) Int() int {
	return int(v)
}

func (v VersionComponent) less(what vComponent) bool {
	if v < what.Component() {
		return true
	}

	return false
}

func (v VersionComponent) is(what vComponent) bool {
	if v == what.Component() {
		return true
	}

	return false
}

func (v VersionComponent) greater(what vComponent) bool {
	if v > what.Component() {
		return true
	}

	return false
}
