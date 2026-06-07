//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.French, "Kirghizistan")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.French, "République kirghize")
	dataKyrgyzstan.RegisterCapital(xlanguage.French, "Bichkek")
}
