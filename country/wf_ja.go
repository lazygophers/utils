//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Japanese, "ウォリス・フツナ")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Japanese, "ウォリス・フツナ")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Japanese, "マタウトゥ")
}
