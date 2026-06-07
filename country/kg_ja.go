//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Japanese, "キルギス")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Japanese, "キルギス共和国")
	dataKyrgyzstan.RegisterCapital(xlanguage.Japanese, "ビシュケク")
}
