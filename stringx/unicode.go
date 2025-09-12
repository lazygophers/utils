package stringx

import (
	"github.com/lazygophers/log"
	"unicode"
)

func AllDigit(s string) bool {
	log.Debugf("AllDigit: checking if all characters in %q are digits", s)
	if s == "" {
		log.Debug("AllDigit: empty string provided, returning true")
		return true
	}
	for _, c := range s {
		if !unicode.IsDigit(c) {
			log.Debugf("AllDigit: found non-digit character %q", c)
			return false
		}
	}
	log.Debug("AllDigit: all characters are digits")
	return true
}

func HasDigit(s string) bool {
	log.Debugf("HasDigit: checking if any character in %q is a digit", s)
	for _, c := range s {
		if unicode.IsDigit(c) {
			log.Debugf("HasDigit: found digit character %q", c)
			return true
		}
	}
	log.Debug("HasDigit: no digit characters found")
	return false
}

func AllLetter(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func HasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func AllSpace(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func HasSpace(s string) bool {
	for _, c := range s {
		if unicode.IsSpace(c) {
			return true
		}
	}
	return false
}

func AllSymbol(s string) bool {
	for _, c := range s {
		if !unicode.IsSymbol(c) {
			return false
		}
	}
	return true
}

func HasSymbol(s string) bool {
	for _, c := range s {
		if unicode.IsSymbol(c) {
			return true
		}
	}
	return false
}

func AllMark(s string) bool {
	for _, c := range s {
		if !unicode.IsMark(c) {
			return false
		}
	}
	return true
}

func HasMark(s string) bool {
	for _, c := range s {
		if unicode.IsMark(c) {
			return true
		}
	}
	return false
}

func AllPunct(s string) bool {
	for _, c := range s {
		if !unicode.IsPunct(c) {
			return false
		}
	}
	return true
}

func HasPunct(s string) bool {
	for _, c := range s {
		if unicode.IsPunct(c) {
			return true
		}
	}
	return false
}

func AllGraphic(s string) bool {
	for _, c := range s {
		if !unicode.IsGraphic(c) {
			return false
		}
	}
	return true
}

func HasGraphic(s string) bool {
	for _, c := range s {
		if unicode.IsGraphic(c) {
			return true
		}
	}
	return false
}

func AllPrint(s string) bool {
	for _, c := range s {
		if !unicode.IsPrint(c) {
			return false
		}
	}
	return true
}

func HasPrint(s string) bool {
	for _, c := range s {
		if unicode.IsPrint(c) {
			return true
		}
	}
	return false
}

func AllControl(s string) bool {
	for _, c := range s {
		if !unicode.IsControl(c) {
			return false
		}
	}
	return true
}

func HasControl(s string) bool {
	for _, c := range s {
		if unicode.IsControl(c) {
			return true
		}
	}
	return false
}

func AllUpper(s string) bool {
	for _, c := range s {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

func HasUpper(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func AllLower(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func HasLower(s string) bool {
	for _, c := range s {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func AllTitle(s string) bool {
	for _, c := range s {
		if !unicode.IsTitle(c) {
			return false
		}
	}
	return true
}

func HasTitle(s string) bool {
	for _, c := range s {
		if unicode.IsTitle(c) {
			return true
		}
	}
	return false
}

func AllLetterOrDigit(s string) bool {
	log.Debugf("AllLetterOrDigit: checking if all characters in %q are letters or digits", s)
	if s == "" {
		log.Debug("AllLetterOrDigit: empty string provided, returning true")
		return true
	}
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			log.Debugf("AllLetterOrDigit: found character %q that is neither letter nor digit", c)
			return false
		}
	}
	log.Debug("AllLetterOrDigit: all characters are letters or digits")
	return true
}

func HasLetterOrDigit(s string) bool {
	log.Debugf("HasLetterOrDigit: checking if any character in %q is a letter or digit", s)
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			log.Debugf("HasLetterOrDigit: found letter or digit character %q", c)
			return true
		}
	}
	log.Debug("HasLetterOrDigit: no letter or digit characters found")
	return false
}
