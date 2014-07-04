/**
 * goversion version management for golang applications
 * Abel Fletcher 2014 abelfletcher@gmail.com
 *
 * This module is used to provide easy version string
 * management. By settings your application's version
 * using goversion.VERSION("0.0.0-rc1"), the
 * version variable will expose the following API
 *
 * Supported formats:
 * 0.0.0      // standard
 * 0.0.0-rc0  // standard with release candidate
 * 0.0.0b0    // standard with beta
 * 0.0.0-r0b0 // standard with release candidate + beta
 *
 * v.String() // returns the full version string
 * v.Major()  // returns the major version component
 * v.Minor()  // returns the minor version component
 * v.Patch()  // returns the patch version component
 * v.IsRc()   // returns release candidate status
 * v.Rc()     // returns rc version
 * v.IsBeta() // returns beta status
 * v.Beta()   // returns beta version
 *
 * Each component exposes the following API
 *
 * c.String()
 *
 * c.Uint8()
 * c.Uint16()
 * c.Uint32()
 * c.Uint64()
 * c.Int8()
 * c.Int16()
 * c.Int32()
 * c.Int64()
 */
package goversion
