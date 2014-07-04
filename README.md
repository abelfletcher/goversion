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

	if v.IsRc() {
		if v.RcIs(4) {
			// do something
		}
	}

	if v.isBeta() {
		if v.BetaIs(22) {
			// do something
		}
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
