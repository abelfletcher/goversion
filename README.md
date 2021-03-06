goversion
=========

goversion provides easy version checking and consistent
serialization for distributed applications.

# Formats

```
0.0.0        // major.minor.patch
0.0.0-rc0    // release candidate
0.0.0b0      // beta
0.0.0-rc0b0  // rc + beta
```

# Example

```
// As application version
var VERSION *goversion.Version

func main() {
	VERSION = goversion.VERSION("0.0.1-rc1")
}
```

```
// As package version
var VERSION *goversion.Version

func init() {
	VERSION = goversion.VERSION("1.0.4")
}
```

```
// As version tracking
func track(version string) {
	v := goversion.VERSION(version)
	
	if v.MajorIs(1) {
		// do something
	} else {
		if v.MinorIs(2) {
			// do something
		} else {
			if v.PatchIs(0) {
				// do something
			}
		}
	}

	// Automatically checks if it is an Rc version
	if v.RcIs(0) {
		// do something
	}

	if v.BetaIs(22) {
		// do something
	}

	if v.LessThanOrEqualTo("1.0.3-rc5") {
		// do something
	} else {
		if v.GreaterThan("0.9.4") {
			// do something else
		}
	}

	// These are the same
	if v.Is("0.0.1") || v.Equals("0.0.1") {
		// do something
	}

	// aliases
	if v.Gte("2.3.4") && v.Lt("2.3.4-rc6b5") {
		// do something
	}
}
```

# Install

```
go get https://github.com/abelfletcher/goversion
```

# Documentation

Full documentation can be generated using godoc.

# Testing

```
go test github.com/abelfletcher/goversion/testing
```
