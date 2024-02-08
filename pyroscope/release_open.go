//go:build release && pyroscope

package pyroscope

func Load(address string) {
	load(address)
}
