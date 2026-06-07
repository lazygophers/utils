//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.Japanese, "インドネシア")
	dataIndonesia.RegisterOfficialName(xlanguage.Japanese, "インドネシア共和国")
	dataIndonesia.RegisterCapital(xlanguage.Japanese, "ジャカルタ")
}
