package test

import (
	"testing"
	"fmt"
	v "github.com/abelfletcher/goversion"
	"strconv"
)

func TestVersion(t *testing.T) {
	newVersion := func(maj int, min int, pat int,
		hrc bool, rc int, hb bool, beta int) *v.Version {
		s := ""
		s += strconv.Itoa(maj)+"."+strconv.Itoa(min)+"."+strconv.Itoa(pat)

		if hrc {
			s += "-rc"+strconv.Itoa(rc)
		}

		if hb {
			s += "b"+strconv.Itoa(beta)
		}

		return v.VERSION(s)
	}

	test := func(v1 *v.Version, v2 *v.Version) bool {
		if !v1.MajorIs(v2.Major().Int()) {
			fmt.Println("Major mismatch")
			return false
		}

		if !v1.MinorIs(v2.Minor().Int()) {
			fmt.Println("Minor mismatch")
			return false
		}

		if !v1.PatchIs(v2.Patch().Int()) {
			fmt.Println("Patch mismatch")
			return false
		}

		if v1.IsRc() != v2.IsRc() {
			fmt.Println("IsRc mismatch")
			return false
		}

		if v1.IsRc() && !v1.RcIs(v2.Rc().Int()) {
			fmt.Println("Rc mismatch")
			return false
		}

		if v1.IsBeta() != v2.IsBeta() {
			fmt.Println("IsBeta mismatch")
			return false
		}

		if v1.IsBeta() && !v1.BetaIs(v2.Beta().Int()) {
			fmt.Println("Beta mismatch")
			return false
		}

		return true
	}

	versions := make(map[string]*v.Version)
	versions["0.0.0"] = newVersion(0, 0, 0, false, 0, false, 0)
	versions["0.0.1"] = newVersion(0, 0, 1, false, 0, false, 0)
	versions["1.3.0"] = newVersion(1, 3, 0, false, 0, false, 0)
	versions["1.4.2-rc1"] = newVersion(1, 4, 2, true, 1, false, 0)
	versions["0.2.21b2"] = newVersion(0, 2, 21, false, 0, true, 2)
	versions["0.3.33-rc2b3"] = newVersion(0, 3, 33, true, 2, true, 3)

	for s, ver := range versions {
		version := v.VERSION(s)

		if !test(version, ver) {
			fmt.Println(s+" not parsed correctly")
			t.Fail()
		} else {
			if version.String() != s {
				fmt.Println(version.String()+" output from String()")
				fmt.Println(s+" output from s")
				fmt.Println("-")
				t.Fail()
			}
		}
	}
}
