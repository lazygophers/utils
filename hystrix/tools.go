package hystrix

import "github.com/lazygophers/utils/randx"

// ProbeWithChance [0, 100]
func ProbeWithChance(percent float64) Probe {
	return func() bool {
		return randx.Booln(percent)
	}
}
