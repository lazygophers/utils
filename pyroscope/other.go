//go:build !release

package pyroscope

func Load(address string) {
	load(address)
}
