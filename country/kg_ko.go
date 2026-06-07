//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Korean, "키르기스스탄")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Korean, "키르기스 공화국")
	dataKyrgyzstan.RegisterCapital(xlanguage.Korean, "비슈케크")
}
