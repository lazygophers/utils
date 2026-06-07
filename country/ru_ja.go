//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.Japanese, "ロシア")
	dataRussia.RegisterOfficialName(xlanguage.Japanese, "ロシア連邦")
	dataRussia.RegisterCapital(xlanguage.Japanese, "モスクワ")
}
