package app

var Organization = "lazygophers"

var Name string

var Version string

type ReleaseType uint8

const (
	Debug ReleaseType = iota
	Test
	Alpha
	Beta
	Release
)

func (p ReleaseType) String() string {
	switch p {
	case Release:
		return "release"
	case Beta:
		return "beta"
	case Alpha:
		return "alpha"
	case Test:
		return "test"
	case Debug:
		fallthrough
	default:
		return "debug"
	}
}

func (p ReleaseType) Debug() string {
	return p.String()
}
