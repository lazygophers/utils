//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Japanese, "エジプト")
	dataEgypt.RegisterOfficialName(xlanguage.Japanese, "エジプト・アラブ共和国")
	dataEgypt.RegisterCapital(xlanguage.Japanese, "カイロ")
}
