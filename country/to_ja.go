//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Japanese, "トンガ")
	dataTonga.RegisterOfficialName(xlanguage.Japanese, "トンガ王国")
	dataTonga.RegisterCapital(xlanguage.Japanese, "ヌクアロファ")
}
