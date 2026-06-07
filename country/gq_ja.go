//go:build (lang_ja || lang_all) && (country_africa || country_all || country_gq || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.Japanese, "赤道ギニア")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.Japanese, "赤道ギニア共和国")
	dataEquatorialGuinea.RegisterCapital(xlanguage.Japanese, "マラボ")
}
