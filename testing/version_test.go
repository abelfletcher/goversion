package test

import (
	"fmt"
	v "github.com/abelfletcher/goversion"
	"strconv"
	"testing"
)

func TestParsing(t *testing.T) {
	newVersion := func(maj int, min int, pat int,
		hrc bool, rc int, hb bool, beta int) *v.Version {
		s := ""
		s += strconv.Itoa(maj) + "." + strconv.Itoa(min) + "." + strconv.Itoa(pat)

		if hrc {
			s += "-rc" + strconv.Itoa(rc)
		}

		if hb {
			s += "b" + strconv.Itoa(beta)
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

	testIs := func(v1 *v.Version, v2 *v.Version) bool {
		if v1.Is(v2.String()) {
			return true
		}

		return false
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
			fmt.Println(s + " not parsed correctly")
			t.Fail()
		} else {
			if version.String() != s {
				fmt.Println(version.String() + " output from String()")
				fmt.Println(s + " output from s")
				fmt.Println("-")
				t.Fail()
			}
		}

		if !testIs(version, ver) {
			fmt.Println("Failed .Is() check")
			t.Fail()
		}

	}
}

func test(cases map[string]map[string]bool, tester func(*v.Version, string, bool)) {
	for verstr, c := range cases {
		version := v.VERSION(verstr)
		for ver, accepted := range c {
			tester(version, ver, accepted)
		}
	}
}

func TestLessThan(t *testing.T) {
	ltVersions := make(map[string]map[string]bool)
	ltVersions["0.2.0"] = make(map[string]bool)
	ltVersions["0.2.0"]["0.2.0"] = false
	ltVersions["0.2.0"]["0.1.9"] = false
	ltVersions["0.2.0"]["0.2.1"] = true
	ltVersions["0.2.0"]["0.2.0b0"] = true
	ltVersions["0.2.0"]["0.2.0-rc1"] = true
	ltVersions["0.2.0"]["0.2.0-rc0b0"] = true
	ltVersions["0.2.0"]["1.0.0"] = true

	ltVersions["4.3.1b0"] = make(map[string]bool)
	ltVersions["4.3.1b0"]["4.3.1b0"] = false
	ltVersions["4.3.1b0"]["4.3.1b1"] = true
	ltVersions["4.3.1b0"]["4.3.1"] = false
	ltVersions["4.3.1b0"]["4.3.1-rc1"] = true
	ltVersions["4.3.1b0"]["4.3.1-rc1b0"] = true
	ltVersions["4.3.1b0"]["5.1.1-rc5"] = true
	ltVersions["4.3.1b0"]["4.3.2"] = true
	ltVersions["4.3.1b0"]["4.3.0-rc5"] = false

	ltVersions["3.5.0-rc3"] = make(map[string]bool)
	ltVersions["3.5.0-rc3"]["3.5.0-rc3"] = false
	ltVersions["3.5.0-rc3"]["3.5.0-rc4"] = true
	ltVersions["3.5.0-rc3"]["3.5.0-rc2"] = false
	ltVersions["3.5.0-rc3"]["3.5.0-rc3b0"] = true
	ltVersions["3.5.0-rc3"]["3.5.1"] = true
	ltVersions["3.5.0-rc3"]["0.10.10-rc15"] = false
	ltVersions["3.5.0-rc3"]["3.5.0-rc2b5"] = false
	ltVersions["3.5.0-rc3"]["3.5.0-rc4b1"] = true

	ltVersions["0.0.5-rc5b3"] = make(map[string]bool)
	ltVersions["0.0.5-rc5b3"]["0.0.5-rc5b3"] = false
	ltVersions["0.0.5-rc5b3"]["0.0.5-rc5b4"] = true
	ltVersions["0.0.5-rc5b3"]["0.0.5-rc5b2"] = false
	ltVersions["0.0.5-rc5b3"]["0.0.5-rc6b0"] = true
	ltVersions["0.0.5-rc5b3"]["0.0.5-rc4b5"] = false
	ltVersions["0.0.5-rc5b3"]["4.3.2-rc3"] = true
	ltVersions["0.0.5-rc5b3"]["0.0.4-rc16"] = false
	ltVersions["0.0.5-rc5b3"]["0.0.5b7"] = false

	test(ltVersions, func(v1 *v.Version, against string, accepted bool) {
		if v1.Lt(against) != accepted {
			t.Fail()
			fmt.Println("Lt check failed. " + v1.String() + ", " + against)
		}
	})
}

func TestLessThanOrEqualTo(t *testing.T) {
	lteVersions := make(map[string]map[string]bool)
	lteVersions["1.3.0"] = make(map[string]bool)
	lteVersions["1.3.0"]["1.3.0"] = true
	lteVersions["1.3.0"]["1.3.1"] = true
	lteVersions["1.3.0"]["1.2.9"] = false
	lteVersions["1.3.0"]["1.3.0b9"] = true
	lteVersions["1.3.0"]["1.3.0-rc1"] = true
	lteVersions["1.3.0"]["1.3.0-rc0b1"] = true
	lteVersions["1.3.0"]["1.2.9-rc9"] = false
	lteVersions["1.3.0"]["5.0.1"] = true

	lteVersions["5.4.3-rc6"] = make(map[string]bool)
	lteVersions["5.4.3-rc6"]["5.4.3-rc6"] = true
	lteVersions["5.4.3-rc6"]["5.4.3-rc5"] = false
	lteVersions["5.4.3-rc6"]["5.4.3-rc7"] = true
	lteVersions["5.4.3-rc6"]["5.4.3-rc6b1"] = true
	lteVersions["5.4.3-rc6"]["5.4.4"] = true
	lteVersions["5.4.3-rc6"]["5.4.3"] = false
	lteVersions["5.4.3-rc6"]["5.4.3b8"] = false

	lteVersions["0.0.8b4"] = make(map[string]bool)
	lteVersions["0.0.8b4"]["0.0.8b4"] = true
	lteVersions["0.0.8b4"]["0.0.8b5"] = true
	lteVersions["0.0.8b4"]["0.0.8b3"] = false
	lteVersions["0.0.8b4"]["0.0.8"] = false
	lteVersions["0.0.8b4"]["0.0.8-rc1"] = true
	lteVersions["0.0.8b4"]["0.0.9"] = true
	lteVersions["0.0.8b4"]["0.0.7"] = false

	lteVersions["1.1.1-rc5b3"] = make(map[string]bool)
	lteVersions["1.1.1-rc5b3"]["1.1.1-rc5b3"] = true
	lteVersions["1.1.1-rc5b3"]["1.1.1-rc5b4"] = true
	lteVersions["1.1.1-rc5b3"]["1.1.1-rc5b2"] = false
	lteVersions["1.1.1-rc5b3"]["1.1.1-rc6b1"] = true
	lteVersions["1.1.1-rc5b3"]["1.1.1-rc4b7"] = false
	lteVersions["1.1.1-rc5b3"]["1.2.1-rc5b3"] = true

	test(lteVersions, func(v1 *v.Version, against string, accepted bool) {
		if v1.Lte(against) != accepted {
			t.Fail()
			fmt.Println("Lte check failed. " + v1.String() + ", " + against)
		}
	})
}

func TestGreaterThan(t *testing.T) {
	gtVersions := make(map[string]map[string]bool)
	gtVersions["2.3.1"] = make(map[string]bool)
	gtVersions["2.3.1"]["2.3.1"] = false
	gtVersions["2.3.1"]["2.3.2"] = false
	gtVersions["2.3.1"]["2.3.0"] = true
	gtVersions["2.3.1"]["2.3.1-rc0"] = false
	gtVersions["2.3.1"]["2.3.1b5"] = false
	gtVersions["2.3.1"]["2.3.1-rc0b5"] = false
	gtVersions["2.3.1"]["2.3.0-rc5"] = true
	gtVersions["2.3.1"]["2.3.0-rc1b9"] = true

	gtVersions["0.4.5b4"] = make(map[string]bool)
	gtVersions["0.4.5b4"]["0.4.5b4"] = false
	gtVersions["0.4.5b4"]["0.4.5b3"] = true
	gtVersions["0.4.5b4"]["0.4.5b5"] = false
	gtVersions["0.4.5b4"]["0.4.4b6"] = true
	gtVersions["0.4.5b4"]["0.4.4-rc1"] = true
	gtVersions["0.4.5b4"]["0.4.5-rc3b1"] = false
	gtVersions["0.4.5b4"]["0.4.5-rc0"] = false

	gtVersions["1.0.3-rc3"] = make(map[string]bool)
	gtVersions["1.0.3-rc3"]["1.0.3-rc3"] = false
	gtVersions["1.0.3-rc3"]["1.0.3-rc2"] = true
	gtVersions["1.0.3-rc3"]["1.0.3-rc4"] = false
	gtVersions["1.0.3-rc3"]["1.0.3-rc3b1"] = false
	gtVersions["1.0.3-rc3"]["1.0.3"] = true
	gtVersions["1.0.3-rc3"]["0.0.3-rc3"] = true
	gtVersions["1.0.3-rc3"]["1.0.3b7"] = true

	gtVersions["2.2.2-rc4b3"] = make(map[string]bool)
	gtVersions["2.2.2-rc4b3"]["2.2.2-rc4b3"] = false
	gtVersions["2.2.2-rc4b3"]["2.2.2-rc5b3"] = false
	gtVersions["2.2.2-rc4b3"]["2.2.2-rc3b3"] = true
	gtVersions["2.2.2-rc4b3"]["2.2.2-rc4b4"] = false
	gtVersions["2.2.2-rc4b3"]["2.2.2-rc4b2"] = true
	gtVersions["2.2.2-rc4b3"]["2.2.2"] = true
	gtVersions["2.2.2-rc4b3"]["2.2.2-rc4"] = true
	gtVersions["2.2.2-rc4b3"]["2.2.2b6"] = true

	test(gtVersions, func(v1 *v.Version, against string, accepted bool) {
		if v1.Gt(against) != accepted {
			t.Fail()
			fmt.Println("Gt check failed. " + v1.String() + ", " + against)
		}
	})
}

func TestGreaterThanOrEqualTo(t *testing.T) {
	gteVersions := make(map[string]map[string]bool)
	gteVersions["5.0.3"] = make(map[string]bool)
	gteVersions["5.0.3"]["5.0.3"] = true
	gteVersions["5.0.3"]["5.0.4"] = false
	gteVersions["5.0.3"]["5.1.0"] = false
	gteVersions["5.0.3"]["2.0.1"] = true
	gteVersions["5.0.3"]["5.0.3b0"] = false
	gteVersions["5.0.3"]["5.0.3-rc1"] = false
	gteVersions["5.0.3"]["5.0.3-rc2b4"] = false
	gteVersions["5.0.3"]["4.1.1"] = true

	gteVersions["3.4.2b3"] = make(map[string]bool)
	gteVersions["3.4.2b3"]["3.4.2b3"] = true
	gteVersions["3.4.2b3"]["3.4.2b1"] = true
	gteVersions["3.4.2b3"]["3.4.2b4"] = false
	gteVersions["3.4.2b3"]["3.4.2-rc1b0"] = false
	gteVersions["3.4.2b3"]["3.4.1-rc5"] = true
	gteVersions["3.4.2b3"]["3.4.2"] = true

	gteVersions["2.4.4-rc1"] = make(map[string]bool)
	gteVersions["2.4.4-rc1"]["2.4.4-rc1"] = true
	gteVersions["2.4.4-rc1"]["2.4.4"] = true
	gteVersions["2.4.4-rc1"]["2.4.4b1"] = true
	gteVersions["2.4.4-rc1"]["2.4.4-rc1b0"] = false
	gteVersions["2.4.4-rc1"]["2.4.3"] = true
	gteVersions["2.4.4-rc1"]["2.4.4-rc2"] = false
	gteVersions["2.4.4-rc1"]["2.4.4-rc0"] = true
	gteVersions["2.4.4-rc1"]["2.4.4-rc0b5"] = true

	gteVersions["0.3.5-rc6b4"] = make(map[string]bool)
	gteVersions["0.3.5-rc6b4"]["0.3.5-rc6b4"] = true
	gteVersions["0.3.5-rc6b4"]["0.3.5-rc6b5"] = false
	gteVersions["0.3.5-rc6b4"]["0.3.5-rc6b3"] = true
	gteVersions["0.3.5-rc6b4"]["0.3.5-rc5b6"] = true
	gteVersions["0.3.5-rc6b4"]["0.3.5-rc7b1"] = false
	gteVersions["0.3.5-rc6b4"]["0.3.5"] = true
	gteVersions["0.3.5-rc6b4"]["0.3.6"] = false

	test(gteVersions, func(v1 *v.Version, against string, accepted bool) {
		if v1.Gte(against) != accepted {
			t.Fail()
			fmt.Println("Gte check failed. " + v1.String() + ", " + against)
		}
	})
}
