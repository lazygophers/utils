//go:build (lang_ja || lang_all) && (country_all || country_oceania || country_polynesia || country_wf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Japanese, "ウォリス・フツナ")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Japanese, "ウォリス・フツナ")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Japanese, "マタウトゥ")
}
