// +build !linux,!darwin,!windows,!openbsd,!freebsd,!netbsd

package defaults

// Sane defaults for the other platforms. The "default" options may be
// may be replaced by the running configuration.
func GetDefaults() platformDefaultParameters {
	return platformDefaultParameters{
		// Admin
		DefaultAdminListen: "tcp://localhost:9001",

		// TUN/TAP
		MaximumIfMTU:     65535,
		DefaultIfMTU:     65535,
		DefaultIfName:    "none",
		DefaultIfTAPMode: false,
	}
}
