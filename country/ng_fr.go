//go:build (lang_fr || lang_all) && (country_africa || country_all || country_ng || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.French, "Nigeria")
	dataNigeria.RegisterOfficialName(xlanguage.French, "République fédérale du Nigeria")
	dataNigeria.RegisterCapital(xlanguage.French, "Abuja")
}
