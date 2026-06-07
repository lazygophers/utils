//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Spanish, "Kirguistán")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Spanish, "República Kirguisa")
	dataKyrgyzstan.RegisterCapital(xlanguage.Spanish, "Biskek")
}
